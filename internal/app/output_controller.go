// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"context"
	"fmt"
	"html"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/therecipe/qt/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

//go:generate qtmoc
type outputController struct {
	core.QObject

	cancelFunc context.CancelFunc
	streamReq  chan *dynamic.Message

	_ int32  `property:"status"`
	_ bool   `property:"running"`
	_ bool   `property:"clientStreaming"`
	_ bool   `property:"bidiStreaming"`
	_ string `property:"output"`
	_ string `property:"stats"`
	_ string `property:"header"`
	_ string `property:"trailer"`

	_ func() `slot:"closeClientStream"`
	_ func() `slot:"cancelRequest"`

	_ func() `constructor:"init"`
}

func (c *outputController) init() {
	c.SetRunning(false)
	c.SetClientStreaming(false)
	c.SetBidiStreaming(false)
	c.SetStatus(-1)

	c.ConnectCloseClientStream(c.closeClientStream)
	c.ConnectCancelRequest(c.cancelRequest)
}

func (c *outputController) clear() {
	c.SetOutput("")
	c.SetStats("")
	c.SetHeader("")
	c.SetTrailer("")
	c.SetStatus(-1)
	c.SetClientStreaming(false)
	c.SetBidiStreaming(false)
}

func (c *outputController) closeClientStream() {
	close(c.streamReq)
}

func (c *outputController) cancelRequest() {
	if c.cancelFunc == nil {
		return
	}
	c.cancelFunc()
}

func (c *outputController) invokeMethod(conn *grpc.ClientConn, md *desc.MethodDescriptor, req *dynamic.Message, meta map[string]string) error {
	if c.IsRunning() && md.IsClientStreaming() {
		c.streamReq <- req
		return nil
	}

	c.clear()

	ctx, cf := context.WithCancel(context.Background())
	c.cancelFunc = cf

	stub := grpcdynamic.NewStub(conn)
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(meta))

	if md.IsClientStreaming() && md.IsServerStreaming() {
		c.SetBidiStreaming(true)
		c.streamReq = make(chan *dynamic.Message)
		go func() {
			stream, err := stub.InvokeRpcBidiStream(ctx, md)
			if err != nil {
				println(err.Error())
				return
			}

			// server streaming
			go func() {
				for {
					_, err := stream.RecvMsg()
					if err != nil {
						break
					}
				}
			}()

			// client streaming
			for r := range c.streamReq {
				if err := stream.SendMsg(r); err != nil {
					if err != io.EOF {
						println(err.Error())
					}
					close(c.streamReq)
				}
			}
			stream.CloseSend()
		}()
		c.streamReq <- req
		return nil
	}

	if md.IsClientStreaming() {
		c.SetClientStreaming(true)
		c.streamReq = make(chan *dynamic.Message)
		go func() {
			stream, err := stub.InvokeRpcClientStream(ctx, md)
			if err != nil {
				println(err.Error())
				return
			}
			for r := range c.streamReq {
				if err := stream.SendMsg(r); err != nil {
					if err != io.EOF {
						println(err.Error())
					}
					close(c.streamReq)
				}
			}
			stream.CloseAndReceive()
		}()
		c.streamReq <- req
		return nil
	}

	if md.IsServerStreaming() {
		go func() {
			stream, err := stub.InvokeRpcServerStream(ctx, md, req)
			if err != nil {
				// TODO: deal with error
				println(err.Error())
				return
			}
			for {
				_, err := stream.RecvMsg()
				if err != nil {
					break
				}
			}
		}()
		return nil
	}

	go func() {
		stub.InvokeRpc(ctx, md, req)
	}()
	return nil
}

// gRPC Stats Handler interface

func (c *outputController) TagRPC(ctx context.Context, _ *stats.RPCTagInfo) context.Context {
	return ctx
}

func (c *outputController) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	statStr := ""
	switch s := stat.(type) {
	case *stats.Begin:
		c.SetRunning(true)
		statStr = formatBegin(s)
	case *stats.OutHeader:
		statStr = formatOutHeader(s)
	case *stats.OutPayload:
		statStr = formatOutPayload(s)
	case *stats.InHeader:
		c.SetHeader(formatMetadata(s.Header))
		statStr = formatInHeader(s)
	case *stats.InTrailer:
		c.SetTrailer(formatMetadata(s.Trailer))
		statStr = formatInTrailer(s)
	case *stats.InPayload:
		dmResp, _ := dynamic.AsDynamicMessage(s.Payload.(proto.Message))
		strResp, _ := marshalTextFormatted(dmResp)
		c.SetOutput(fmt.Sprintf("%s<p>%s</p>", c.Output(), string(strResp)))
		statStr = formatInPayload(s)
	case *stats.End:
		c.SetRunning(false)
		if s.Error == nil {
			c.SetStatus(0)
		}
		if s.Error != nil {
			dmErr, _ := dynamic.AsDynamicMessage(status.Convert(s.Error).Proto())
			strErr, _ := marshalTextFormatted(dmErr)
			c.SetStatus(int(status.Code(s.Error)))
			c.SetOutput(fmt.Sprintf("%s%s<br/>", c.Output(), string(strErr)))
		}
		statStr = formatEnd(s)
	}
	c.SetStats(fmt.Sprintf("%s%s", c.Stats(), statStr))
}

func (c *outputController) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	return ctx
}

func (c *outputController) HandleConn(ctx context.Context, stat stats.ConnStats) {}

func formatMetadata(md metadata.MD) string {
	keys := make([]string, 0, len(md))
	for k := range md {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString("<p>")
	for _, k := range keys {
		sb.WriteString("<span>")
		sb.WriteString(k)
		sb.WriteString(": </span>")
		sb.WriteString(strings.Join(md.Get(k), ", "))
		sb.WriteString("<br/>")
	}
	sb.WriteString("</p>")
	return sb.String()
}

func formatBegin(s *stats.Begin) string {
	var sb strings.Builder
	sb.WriteString("<p>")
	sb.WriteString("<b>Begin</b>")
	sb.WriteString("<br/>")
	sb.WriteString("Begin Time: ")
	sb.WriteString(s.BeginTime.String())
	sb.WriteString("<br/>")
	sb.WriteString("Fail Fast: ")
	sb.WriteString(strconv.FormatBool(s.FailFast))
	sb.WriteString("</p>")
	return sb.String()
}

func formatOutHeader(s *stats.OutHeader) string {
	compression := s.Compression
	if compression == "" {
		compression = "nil"
	}
	var sb strings.Builder
	sb.WriteString("<p class='yellow'>")
	sb.WriteString("» <b>Out Header</b>")
	sb.WriteString("<br/>")
	sb.WriteString("» Compression: ")
	sb.WriteString(compression)
	sb.WriteString("<br/>")
	sb.WriteString("» Header: ")
	sb.WriteString(fmt.Sprintf("%+v", s.Header))
	sb.WriteString("<br/>")
	sb.WriteString("» Full Method: ")
	sb.WriteString(s.FullMethod)
	sb.WriteString("<br/>")
	sb.WriteString("» Remote Address: ")
	sb.WriteString(s.RemoteAddr.String())
	sb.WriteString("<br/>")
	sb.WriteString("» Local Address: ")
	sb.WriteString(s.LocalAddr.String())
	sb.WriteString("<p>")
	return sb.String()
}

func formatOutPayload(s *stats.OutPayload) string {
	p := proto.MarshalTextString(s.Payload.(proto.Message))
	var sb strings.Builder
	sb.WriteString("<p class='yellow'>")
	sb.WriteString("» <b>Out Payload</b>")
	sb.WriteString("<br/>")
	sb.WriteString("» Payload: ")
	sb.WriteString(html.EscapeString(p))
	sb.WriteString("<br/>")
	sb.WriteString("» Data: ")
	sb.WriteString(fmt.Sprintf("%+v", s.Data))
	sb.WriteString("<br/>")
	sb.WriteString("» Length: ")
	sb.WriteString(strconv.Itoa(s.Length))
	sb.WriteString("<br/>")
	sb.WriteString("» Wire Length: ")
	sb.WriteString(strconv.Itoa(s.WireLength))
	sb.WriteString("<br/>")
	sb.WriteString("» Sent Time: ")
	sb.WriteString(s.SentTime.String())
	sb.WriteString("<p>")
	return sb.String()
}

func formatInHeader(s *stats.InHeader) string {
	compression := s.Compression
	if compression == "" {
		compression = "nil"
	}
	var sb strings.Builder
	sb.WriteString("<p class='green'>")
	sb.WriteString("« <b>In Header</b>")
	sb.WriteString("<br/>")
	sb.WriteString("« Wire Length: ")
	sb.WriteString(strconv.Itoa(s.WireLength))
	sb.WriteString("<br/>")
	sb.WriteString("« Compression: ")
	sb.WriteString(compression)
	sb.WriteString("<br/>")
	sb.WriteString("« Header: ")
	sb.WriteString(fmt.Sprintf("%+v", s.Header))
	sb.WriteString("<p>")
	return sb.String()
}

func formatInTrailer(s *stats.InTrailer) string {
	var sb strings.Builder
	sb.WriteString("<p class='green'>")
	sb.WriteString("« <b>In Trailer</b>")
	sb.WriteString("<br/>")
	sb.WriteString("« Wire Length: ")
	sb.WriteString(strconv.Itoa(s.WireLength))
	sb.WriteString("<br/>")
	sb.WriteString("« Trailer: ")
	sb.WriteString(fmt.Sprintf("%+v", s.Trailer))
	sb.WriteString("<p>")
	return sb.String()
}

func formatInPayload(s *stats.InPayload) string {
	p := proto.MarshalTextString(s.Payload.(proto.Message))
	var sb strings.Builder
	sb.WriteString("<p class='green'>")
	sb.WriteString("« <b>In Payload</b>")
	sb.WriteString("<br/>")
	sb.WriteString("« Payload: ")
	sb.WriteString(html.EscapeString(p))
	sb.WriteString("<br/>")
	sb.WriteString("« Data: ")
	sb.WriteString(fmt.Sprintf("%+v", s.Data))
	sb.WriteString("<br/>")
	sb.WriteString("« Length: ")
	sb.WriteString(strconv.Itoa(s.Length))
	sb.WriteString("<br/>")
	sb.WriteString("« Wire Length: ")
	sb.WriteString(strconv.Itoa(s.WireLength))
	sb.WriteString("<br/>")
	sb.WriteString("« Recived Time: ")
	sb.WriteString(s.RecvTime.String())
	sb.WriteString("<p>")
	return sb.String()
}

func formatEnd(s *stats.End) string {
	var sb strings.Builder
	sb.WriteString("<p>")
	sb.WriteString("<b>End</b>")
	sb.WriteString("<br/>")
	sb.WriteString("Begin Time: ")
	sb.WriteString(s.BeginTime.String())
	sb.WriteString("<br/>")
	sb.WriteString("End Time: ")
	sb.WriteString(s.EndTime.String())
	sb.WriteString("<br/>")
	if s.Error != nil {
		sb.WriteString("<span class='red'>")
		sb.WriteString("Error: ")
		sb.WriteString(s.Error.Error())
		sb.WriteString("</span>")
	}
	sb.WriteString("</p>")
	return sb.String()
}

// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"context"
	"errors"
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

	_ int32  `property:"status"`
	_ string `property:"output"`
	_ string `property:"stats"`
	_ string `property:"header"`
	_ string `property:"trailer"`

	_ func() `constructor:"init"`
}

func (c *outputController) init() {
	c.SetStatus(-1)
}

func (c *outputController) clear() {
	c.SetOutput("")
	c.SetStats("")
	c.SetHeader("")
	c.SetTrailer("")
	c.SetStatus(-1)
}

func (c *outputController) invokeMethod(conn *grpc.ClientConn, md *desc.MethodDescriptor, req *dynamic.Message, meta map[string]string) error {
	c.clear()

	stub := grpcdynamic.NewStub(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(meta))

	if md.IsClientStreaming() && md.IsServerStreaming() {
		return errors.New("Bidirectional streaming not supported yet")
	}

	if md.IsClientStreaming() {
		return errors.New("Client streaming not supported yet")
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
				resp, err := stream.RecvMsg()
				c.processResponse(resp, err, true)
				if err != nil {
					break
				}
			}
		}()
		return nil
	}

	go func() {
		resp, err := stub.InvokeRpc(ctx, md, req)
		c.processResponse(resp, err, false)
	}()
	return nil
}

func (c *outputController) processResponse(resp proto.Message, err error, isStream bool) {
	if err != nil {
		if isStream && err == io.EOF {
			MainThread.Run(func() {
				c.SetStatus(0)
			})
			return
		}
		dmErr, _ := dynamic.AsDynamicMessage(status.Convert(err).Proto())
		strErr, _ := dmErr.MarshalTextIndent()
		MainThread.Run(func() {
			c.SetStatus(int(status.Code(err)))
			// TODO: we should stream the data to the UI so we can use the TextEdit append
			// function, which would have better performance than replacing the whole text
			c.SetOutput(fmt.Sprintf("%s%s\n", c.Output(), string(strErr)))
		})
		return
	}

	dmResp, _ := dynamic.AsDynamicMessage(resp)
	strResp, _ := dmResp.MarshalTextIndent()

	MainThread.Run(func() {
		if !isStream {
			c.SetStatus(0)
		}
		c.SetOutput(fmt.Sprintf("%s%s\n", c.Output(), string(strResp)))
	})
}

// gRPC Stats Handler interface

func (c *outputController) TagRPC(ctx context.Context, _ *stats.RPCTagInfo) context.Context {
	return ctx
}

func (c *outputController) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	statStr := ""
	switch s := stat.(type) {
	case *stats.Begin:
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
		statStr = formatInPayload(s)
	case *stats.End:
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

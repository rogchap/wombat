// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"context"
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/therecipe/qt/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"rogchap.com/wombat/internal/model"
)

//go:generate qtmoc
type outputController struct {
	core.QObject

	_ int32  `property:"status"`
	_ string `property:"output"` // Temp property, should be a list model

	_ *model.KeyvalList `property:"headers"`
	_ *model.KeyvalList `property:"trailers"`

	_ func() `constructor:"init"`
}

func (c *outputController) init() {
	c.SetStatus(-1)
	c.SetHeaders(model.NewKeyvalList(nil))
	c.SetTrailers(model.NewKeyvalList(nil))
}

func (c *outputController) clear() {
	c.SetOutput("")
	c.SetStatus(-1)
	c.Headers().UpdateList(nil)
	c.Trailers().UpdateList(nil)
}

func (c *outputController) invokeMethod(conn *grpc.ClientConn, md *desc.MethodDescriptor, req *dynamic.Message, meta map[string]string) {
	c.clear()

	stub := grpcdynamic.NewStub(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(meta))

	if md.IsClientStreaming() && md.IsServerStreaming() {
		println("Bidirectional streaming not supported yet")
		return
	}

	if md.IsClientStreaming() {
		println("Client streaming not supported yet")
		return
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
		return
	}

	go func() {
		resp, err := stub.InvokeRpc(ctx, md, req)
		c.processResponse(resp, err, false)
	}()

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
	fmt.Printf("%T: %+[1]v\n\n", stat)
	switch s := stat.(type) {
	case *stats.InHeader:
		c.Headers().UpdateList(model.MapMetadata(s.Header))
	case *stats.InTrailer:
		c.Trailers().SetList(model.MapMetadata(s.Trailer))
	}
}

func (c *outputController) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	return ctx
}

func (c *outputController) HandleConn(ctx context.Context, stat stats.ConnStats) {
	fmt.Printf("%T: %+[1]v\n\n", stat)
}

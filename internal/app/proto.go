package app

import (
	"context"
	"errors"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ctxInternalKey struct{}

func protoFilesFromReflectionAPI(conn *grpc.ClientConn, md metadata.MD) (*protoregistry.Files, error) {
	if conn == nil {
		return nil, errors.New("app: no connection to a grpc server available")
	}

	stub := rpb.NewServerReflectionClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx = context.WithValue(ctx, ctxInternalKey{}, struct{}{})
	client := grpcreflect.NewClient(ctx, stub)
	defer client.Reset()

	services, err := client.ListServices()
	if err != nil {
		return nil, err
	}

	seen := make(map[string]struct{})
	fdset := &descriptorpb.FileDescriptorSet{}

	for _, srv := range services {
		fd, err := client.FileContainingSymbol(srv)
		if err != nil {
			return nil, err
		}
		fdset.File = append(fdset.File, walkFileDescriptors(seen, fd)...)
	}

	return protodesc.NewFiles(fdset)
}

func protoFilesFromDisk(importPaths, filenames []string) (*protoregistry.Files, error) {
	if len(filenames) == 0 {
		return nil, errors.New("app: no *.proto files found")
	}

	f, err := protoparse.ResolveFilenames(importPaths, filenames...)
	if err != nil {
		return nil, err
	}

	parser := protoparse.Parser{
		ImportPaths:      importPaths,
		InferImportPaths: len(importPaths) == 0,
	}

	fds, err := parser.ParseFiles(f...)
	if err != nil {
		return nil, err
	}

	fdset := &descriptorpb.FileDescriptorSet{}
	seen := make(map[string]struct{})

	for _, fd := range fds {
		fdset.File = append(fdset.File, walkFileDescriptors(seen, fd)...)
	}

	return protodesc.NewFiles(fdset)
}

func walkFileDescriptors(seen map[string]struct{}, fd *desc.FileDescriptor) []*descriptorpb.FileDescriptorProto {
	fds := []*descriptorpb.FileDescriptorProto{}

	if _, ok := seen[fd.GetName()]; ok {
		return fds
	}
	seen[fd.GetName()] = struct{}{}
	fds = append(fds, fd.AsFileDescriptorProto())

	for _, dep := range fd.GetDependencies() {
		deps := walkFileDescriptors(seen, dep)
		fds = append(fds, deps...)
	}

	return fds
}

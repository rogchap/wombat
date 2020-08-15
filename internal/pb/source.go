// Copyright 2020 Rogchap. All Rights Reserved.

package pb

import (
	"context"
	"errors"
	"fmt"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type ctxInternal struct{}

type Source interface {
	Services() []string
	Methods() map[string][]*desc.MethodDescriptor
	GetMethodDesc(srv, name string) *desc.MethodDescriptor
}

type source struct {
	services []string
	methods  map[string][]*desc.MethodDescriptor
}

func (s *source) Services() []string {
	return s.services
}

func (s *source) Methods() map[string][]*desc.MethodDescriptor {
	return s.methods
}

func (s *source) GetMethodDesc(srv, name string) *desc.MethodDescriptor {
	methods := s.methods[srv]
	if methods == nil {
		return nil
	}

	for _, md := range methods {
		if md.GetName() == name {
			return md
		}
	}

	return nil
}

func GetSourceFromProtoFiles(importPaths, protoPaths []string) (Source, error) {
	filenames, err := protoparse.ResolveFilenames(importPaths, protoPaths...)
	if err != nil {
		return nil, fmt.Errorf("pb: failed to resolve filenames: %v", err)
	}
	parser := protoparse.Parser{
		ImportPaths:      importPaths,
		InferImportPaths: len(importPaths) == 0,
	}
	fds, err := parser.ParseFiles(filenames...)
	if err != nil {
		return nil, fmt.Errorf("pb: failed to parse proto files: %v", err)
	}

	var services []string
	methods := make(map[string][]*desc.MethodDescriptor)
	for _, fd := range fds {
		for _, srv := range fd.GetServices() {
			srvName := srv.GetFullyQualifiedName()
			services = append(services, srvName)
			var ms []*desc.MethodDescriptor
			for _, m := range srv.GetMethods() {
				ms = append(ms, m)
			}
			methods[srvName] = ms
		}
	}

	return &source{
		// files:    fds,
		services: services,
		methods:  methods,
	}, nil
}

func GetSourceFromReflectionAPI(conn *grpc.ClientConn) (Source, error) {
	if conn == nil {
		return nil, errors.New("pb: no connection available")
	}

	stub := rpb.NewServerReflectionClient(conn)
	ctx := context.WithValue(context.Background(), "ctxInternal", struct{}{})
	client := grpcreflect.NewClient(ctx, stub)
	defer client.Reset()

	services, err := client.ListServices()
	if err != nil {
		return nil, err
	}

	methods := make(map[string][]*desc.MethodDescriptor)
	for _, srv := range services {
		sd, err := client.ResolveService(srv)

		if err != nil {
			return nil, err
		}
		var ms []*desc.MethodDescriptor
		for _, md := range sd.GetMethods() {
			ms = append(ms, md)
		}
		methods[srv] = ms

	}

	return &source{
		services: services,
		methods:  methods,
	}, nil

}

package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc --go_out=:. --go-grpc_out=:. route_guide.proto foobar.proto

type server struct {
	UnimplementedRouteGuideServer
	UnimplementedFoobarServer

	savedFeatures []*Feature // read-only after initialized

	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*RouteNote
}

func newServer() *server {
	s := &server{routeNotes: make(map[string][]*RouteNote)}
	json.Unmarshal(exampleData, &s.savedFeatures)

	return s
}

// GetFeature returns the feature at the given point.
func (s *server) GetFeature(ctx context.Context, point *Point) (*Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	// No feature was found, return an unnamed feature
	return &Feature{Location: point}, nil
}

// Serve stats serving a gRPC server that is used for testing
func Serve() {
	lis, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		fmt.Fprintf(os.Stderr, "server: failed to create listener: %v", err)
	}

	s := newServer()
	gs := grpc.NewServer()
	RegisterRouteGuideServer(gs, s)
	RegisterFoobarServer(gs, s)
	reflection.Register(gs)
	gs.Serve(lis)
}

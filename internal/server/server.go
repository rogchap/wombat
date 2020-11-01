package server

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
)

func init() {
	rand.Seed(time.Now().Unix())
}

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

func inRange(point *Point, rect *Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (s *server) ListFeatures(rect *Rectangle, stream RouteGuide_ListFeaturesServer) error {
	for _, feature := range s.savedFeatures {
		if inRange(feature.Location, rect) {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
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

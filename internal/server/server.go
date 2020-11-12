package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// calcDistance calculates the distance between two points using the "haversine" formula.
// The formula is based on http://mathforum.org/library/drmath/view/51879.html.
func calcDistance(p1 *Point, p2 *Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points,  number of known features visited, total distance traveled, and
// total time spent.
func (s *server) RecordRoute(stream RouteGuide_RecordRouteServer) error {
	var pointCount, featureCount, distance int32
	var lastPoint *Point
	startTime := time.Now()
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		pointCount++
		for _, feature := range s.savedFeatures {
			if proto.Equal(feature.Location, point) {
				featureCount++
			}
		}
		if lastPoint != nil {
			distance += calcDistance(lastPoint, point)
		}
		lastPoint = point
	}
}

// Serve stats serving a gRPC server that is used for testing
func Serve() {
	lis, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		fmt.Fprintf(os.Stderr, "server: failed to create listener: %v", err)
	}

	e, _ := protojson.Marshal(&WellKnownRequest{Timestamp: timestamppb.Now()})
	fmt.Printf("string(e) = %+v\n", string(e))

	s := newServer()
	gs := grpc.NewServer()
	RegisterRouteGuideServer(gs, s)
	RegisterFoobarServer(gs, s)
	reflection.Register(gs)
	gs.Serve(lis)
}

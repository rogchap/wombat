package app

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type cyclicDetector struct {
    seen  map[protoreflect.FullName]struct{}
    graph []string
    depth int // Added to track the depth of the cycle
}

func (d *cyclicDetector) detect(md protoreflect.MessageDescriptor) error {
    if d.seen == nil {
        d.seen = make(map[protoreflect.FullName]struct{})
        d.depth = 0 // Reset depth when starting a new detection
    }

    n, f := md.Name(), md.FullName()
    if _, cyclic := d.seen[f]; cyclic {
        d.depth++ // Increment depth if a cycle is detected
        d.graph = append(d.graph, string(n))
        if d.depth > 7 { // Check if the depth exceeds 7
            return fmt.Errorf("unable to parse proto descriptors: cyclic data detected beyond allowed depth: %s", strings.Join(d.graph, " → "))
        }
    } else {
        d.seen[f] = struct{}{}
        d.graph = append(d.graph, string(n))
        d.depth = 0 // Reset depth if no cycle is detected
    }

    return nil
}

func (d *cyclicDetector) reset() {
    d.seen = nil
    d.graph = nil
    d.depth = 0 // Ensure depth is reset here as well
}
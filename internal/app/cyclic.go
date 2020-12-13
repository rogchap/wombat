package app

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type cyclicDetector struct {
	seen  map[protoreflect.FullName]struct{}
	graph []string
}

func (d *cyclicDetector) detect(md protoreflect.MessageDescriptor) error {
	if d.seen == nil {
		d.seen = make(map[protoreflect.FullName]struct{})
	}

	n, f := md.Name(), md.FullName()
	if _, cyclic := d.seen[f]; cyclic {
		d.graph = append(d.graph, string(n))
		return fmt.Errorf("unable to parse proto descriptors: cyclic data detected: %s", strings.Join(d.graph, " → "))
	}
	d.seen[f] = struct{}{}
	d.graph = append(d.graph, string(n))

	return nil
}

func (d *cyclicDetector) reset() {
	d.seen = nil
	d.graph = nil
}

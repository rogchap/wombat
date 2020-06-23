// Copyright 2020 Rogchap. All Rights Reserved.

package pb_test

import (
	"testing"

	"rogchap.com/courier/internal/pb"
)

func TestGetSourceFromProtoFiles(t *testing.T) {
	_, err := pb.GetSourceFromProtoFiles(nil, []string{"../../example/route_guide.proto"})
	if err != nil {
		t.Error(err)
	}
}

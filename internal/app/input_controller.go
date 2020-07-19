// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"errors"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"rogchap.com/courier/internal/model"
	"rogchap.com/courier/internal/pb"
)

//go:generate qtmoc
type inputController struct {
	core.QObject

	pbSource pb.Source

	_ func() `constructor:"init"`

	_ *model.StringList `property:"serviceListModel"`
	_ *model.StringList `property:"methodListModel"`
	_ *model.Message    `property:"requestModel"`

	_ func(service string)         `slot:"serviceChanged"`
	_ func(service, method string) `slot:"methodChanged"`
}

func (c *inputController) init() {
	c.SetServiceListModel(model.NewStringList(nil))
	c.SetMethodListModel(model.NewStringList(nil))
	c.SetRequestModel(model.NewMessage(nil))

	c.ConnectServiceChanged(c.serviceChanged)
	c.ConnectMethodChanged(c.methodChanged)
}

func (c *inputController) processProtos(imports, protos []string) error {
	if len(protos) == 0 {
		return errors.New("no *.proto files to process")
	}
	if len(imports) == 0 {
		// optomistacally try and use a import path
		imports = append(imports, filepath.Dir(protos[0]))
	}

	var err error
	c.pbSource, err = pb.GetSourceFromProtoFiles(imports, protos)
	if err != nil {
		return err
	}

	services := c.pbSource.Services()
	if len(services) == 0 {
		return errors.New("no gRPC services found in proto files")
	}

	c.ServiceListModel().SetStringList(services)
	c.serviceChanged(services[0])
	return nil
}

func (c *inputController) serviceChanged(service string) {
	methods := c.pbSource.Methods()

	srvMethods, ok := methods[service]
	if !ok {
		return
	}
	var methodStrs []string
	for _, m := range srvMethods {
		methodStrs = append(methodStrs, m.GetName())
	}

	c.MethodListModel().SetStringList(methodStrs)
	c.methodChanged(service, methodStrs[0])
}

func (c *inputController) methodChanged(service, method string) {
	md := c.pbSource.GetMethodDesc(service, method)
	if md == nil {
		return
	}

	c.SetRequestModel(model.MapMessage(md.GetInputType()))
}

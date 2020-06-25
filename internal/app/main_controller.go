// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"

	"rogchap.com/courier/internal/model"
	"rogchap.com/courier/internal/pb"
)

type mainController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ *model.ServiceList `property:"serviceList"`
	_ *model.MethodList  `property:"methodList"`

	_ func(path string) `slot:"processProtos"`
}

func (c *mainController) init() {

	c.SetServiceList(model.NewServiceList(nil))
	c.SetMethodList(model.NewMethodList(nil))

	c.ConnectProcessProtos(c.processProtos)
}

func (c *mainController) processProtos(path string) {
	var protoFiles []string
	filepath.Walk(path[7:], func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})

	if len(protoFiles) == 0 {
		return
		// TODO: Show error to user that there is no proto files found
	}

	source, err := pb.GetSourceFromProtoFiles([]string{path[7:]}, protoFiles)
	if err != nil {
		println(err.Error())
		return
	}

	services := source.Services()
	if len(services) == 0 {
		// TODO: Show error that there are no servcies found
		return
	}
	c.ServiceList().SetStringList(services)
	c.MethodList().SetStringList(source.Methods()[services[0]])
}

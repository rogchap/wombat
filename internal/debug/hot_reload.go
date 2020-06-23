// Copyright 2020 Rogchap. All Rights Reserved.

package debug

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

type Reloader interface {
	Close()
}

type hotreloader struct {
	*core.QFileSystemWatcher
}

func (h *hotreloader) Close() {
	h.DisconnectDirectoryChanged()
}

// HotReloader will watch the ./qml folder recursively and re-load the main.qml file
func HotReloader(engine *qml.QQmlApplicationEngine) Reloader {
	qmlRoot := filepath.Join(".", "qml")
	watcher := core.NewQFileSystemWatcher2([]string{qmlRoot}, nil)
	watcher.ConnectDirectoryChanged(func(p string) {
		for _, o := range engine.RootObjects() {
			if o.IsWindowType() {
				// This will still leave to window in RootObjects; how can we remove?
				gui.NewQWindowFromPointer(o.Pointer()).Close()
			}
		}
		engine.ClearComponentCache()
		engine.Load(core.NewQUrl3(filepath.Join(qmlRoot, "main.qml"), 0))
	})

	filepath.Walk(qmlRoot, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsDir() {
			watcher.AddPath(path)
		}
		return nil
	})

	return &hotreloader{watcher}
}

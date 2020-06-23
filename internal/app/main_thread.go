// Copyright 2020 Rogchap. All Rights Reserved.

package app

import "github.com/therecipe/qt/core"

// MainThread allows is to use a signal to run code within a goroutine back on the main thread
var MainThread = NewMainThreadRunner(nil)

type mainThreadRunner struct {
	core.QObject

	_ func(f func()) `signal:"run,auto"`
}

func (*mainThreadRunner) run(f func()) {
	f()
}

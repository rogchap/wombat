package app

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"
)

type api struct {
	runtime       *wails.Runtime
	logger        *logger.CustomLogger
	client        *client
	cancelCtxFunc context.CancelFunc
}

// WailsInit is the init fuction for the wails runtime
func (a *api) WailsInit(runtime *wails.Runtime) error {
	a.runtime = runtime
	a.logger = runtime.Log.New("API")
	return nil
}

func (a *api) monitorStateChanges(ctx context.Context) {
	for {
		if a.client == nil || a.client.conn == nil {
			continue
		}
		state := a.client.conn.GetState()
		a.runtime.Events.Emit(eventClientStateChanged, state.String())
		if ok := a.client.conn.WaitForStateChange(ctx, state); !ok {
			a.logger.Debug("ending monitoring of state changes")
			return
		}
	}
}

// Connect will attempt to connect a grpc server and parse any proto files
func (a *api) Connect(data interface{}) error {
	var opts options
	if err := mapstructure.Decode(data, &opts); err != nil {
		return err
	}

	if a.client != nil {
		if err := a.client.close(); err != nil {
			return fmt.Errorf("app: failed to close previous connection: %v", err)
		}
	}

	if a.cancelCtxFunc != nil {
		a.cancelCtxFunc()
	}

	a.client = &client{}
	if err := a.client.connect(opts); err != nil {
		return fmt.Errorf("app: failed to connect to server: %v", err)
	}

	a.runtime.Events.Emit(eventClientConnected, opts.Addr)

	ctx := context.Background()
	ctx, a.cancelCtxFunc = context.WithCancel(ctx)
	go a.monitorStateChanges(ctx)

	return nil
}

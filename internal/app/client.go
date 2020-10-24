package app

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type client struct {
	conn *grpc.ClientConn
}

type options struct {
	Addr    string `json:"addr"`
	Reflect bool   `json:"reflect"`

	Insecure  bool `json:"insecure"`
	Plaintext bool `json:"plaintext"`
}

func (c *client) connect(o options) error {
	errc := make(chan error, 1)
	go func() {
		opts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
			// grpc.WithStatsHandler(c),
		}

		// TODO: wombat user agent

		if !o.Plaintext {
			// TODO: tls
		}

		if o.Plaintext {
			opts = append(opts, grpc.WithInsecure())
		}

		dialer := func(ctx context.Context, addr string) (net.Conn, error) {
			d := &net.Dialer{}
			return d.DialContext(ctx, "tcp", addr)
		}
		opts = append(opts, grpc.WithContextDialer(dialer))

		var err error
		c.conn, err = grpc.Dial(o.Addr, opts...)
		if err != nil {
			errc <- err

			return
		}
		close(errc)
	}()

	if err := <-errc; err != nil {
		return err
	}
	return nil
}

func (c *client) close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

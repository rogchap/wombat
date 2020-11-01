package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/stats"
	"google.golang.org/protobuf/proto"
)

type client struct {
	conn *grpc.ClientConn
}

type transportCreds struct {
	credentials.TransportCredentials
	errc chan<- error
}

func (t *transportCreds) ClientHandshake(ctx context.Context, addr string, in net.Conn) (net.Conn, credentials.AuthInfo, error) {
	out, auth, err := t.TransportCredentials.ClientHandshake(ctx, addr, in)
	if err != nil {
		t.errc <- err
	}
	return out, auth, err
}

func (c *client) connect(o options, h stats.Handler) error {
	errc := make(chan error, 1)
	go func() {
		opts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithStatsHandler(h),
		}

		// TODO: wombat user agent

		if !o.Plaintext {
			var tlsCfg tls.Config
			tlsCfg.InsecureSkipVerify = o.Insecure

			if o.Clientcert != "" {
				cert, err := tls.X509KeyPair([]byte(o.Clientcert), []byte(o.Clientkey))
				if err != nil {
					errc <- err
					return
				}
				tlsCfg.Certificates = []tls.Certificate{cert}
			}

			var err error
			tlsCfg.RootCAs, err = x509.SystemCertPool()
			if err != nil {
				tlsCfg.RootCAs = x509.NewCertPool()
			}
			if o.Rootca != "" {
				tlsCfg.RootCAs.AppendCertsFromPEM([]byte(o.Rootca))
			}
			creds := &transportCreds{
				credentials.NewTLS(&tlsCfg),
				errc,
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
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

func (c *client) invoke(ctx context.Context, method string, req, resp proto.Message) error {
	if c.conn == nil {
		return errors.New("app: no connection available")
	}

	return c.conn.Invoke(ctx, method, req, resp)
}

func (c *client) close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

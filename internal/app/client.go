package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/stats"
	"google.golang.org/protobuf/proto"
)

var errNoConn = errors.New("app: no connection available")

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
			grpc.WithUserAgent(fmt.Sprintf("%s/%s", appname, semver)),
		}

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
		}

		if o.Plaintext {
			opts = append(opts, grpc.WithInsecure())
		}

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
		return errNoConn
	}

	return c.conn.Invoke(ctx, method, req, resp)
}

func (c *client) invokeServerStream(ctx context.Context, method string, req proto.Message) (grpc.ClientStream, error) {
	if c.conn == nil {
		return nil, errNoConn
	}
	sd := &grpc.StreamDesc{
		StreamName:    method,
		ClientStreams: false,
		ServerStreams: true,
	}
	ctx, cancel := context.WithCancel(ctx)
	_ = cancel // avoid go vet error
	s, err := c.conn.NewStream(ctx, sd, method)
	if err != nil {
		return nil, err
	}
	if err := s.SendMsg(req); err != nil {
		cancel()
		return nil, err
	}
	if err := s.CloseSend(); err != nil {
		cancel()
		return nil, err
	}
	return s, nil
}

func (c *client) invokeClientStream(ctx context.Context, method string) (grpc.ClientStream, error) {
	if c.conn == nil {
		return nil, errNoConn
	}
	sd := &grpc.StreamDesc{
		StreamName:    method,
		ClientStreams: true,
		ServerStreams: false,
	}
	return c.conn.NewStream(ctx, sd, method)
}

func (c *client) invokeBidiStream(ctx context.Context, method string) (grpc.ClientStream, error) {
	if c.conn == nil {
		return nil, errNoConn
	}
	sd := &grpc.StreamDesc{
		StreamName:    method,
		ClientStreams: true,
		ServerStreams: true,
	}
	return c.conn.NewStream(ctx, sd, method)
}

func (c *client) close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

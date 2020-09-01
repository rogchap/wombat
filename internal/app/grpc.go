// Copyright 2020 Rogchap. All Rights Reserved.

package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/stats"

	"rogchap.com/wombat/internal/model"
)

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

// BlockDIal isa wrapper to the regular grpc dial, but will block until the connection is made or an error occurs.
func BlockDial(addr string, opts *model.WorkspaceOptions, statHandler stats.Handler) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	errc := make(chan error)

	go func() {
		dopts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithStatsHandler(statHandler),
			grpc.WithUserAgent(fmt.Sprintf("%s/%s", appname, semver)),
		}

		if !opts.IsPlaintext() {
			var tlsCfg tls.Config
			tlsCfg.InsecureSkipVerify = opts.IsInsecure()

			if opts.Clientcert() != "" {
				cert, err := tls.X509KeyPair([]byte(opts.Clientcert()), []byte(opts.Clientkey()))
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
			if opts.Rootca() != "" {
				tlsCfg.RootCAs.AppendCertsFromPEM([]byte(opts.Rootca()))
			}
			creds := &transportCreds{
				credentials.NewTLS(&tlsCfg),
				errc,
			}
			dopts = append(dopts, grpc.WithTransportCredentials(creds))
		}

		if opts.IsPlaintext() {
			dopts = append(dopts, grpc.WithInsecure())
		}

		ctxDialer := func(ctx context.Context, addr string) (net.Conn, error) {
			d := &net.Dialer{}
			return d.DialContext(ctx, "tcp", addr)
		}
		dopts = append(dopts, grpc.WithContextDialer(ctxDialer))

		var err error
		conn, err = grpc.Dial(addr, dopts...)
		if err != nil {
			errc <- err
			return
		}
		close(errc)
	}()

	select {
	case err := <-errc:
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
}

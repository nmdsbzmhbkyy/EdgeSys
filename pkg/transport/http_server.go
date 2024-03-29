package transport

import (
	"EdgeSys/pkg/global"
	"context"
	"mod.miligc.com/edge-common/business-common/business/pkg"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type HttpServer struct {
	Addr string
	srv  *http.Server

	Container *restful.Container
}

func NewHttpServer(addr string) *HttpServer {
	c := restful.NewContainer()
	c.EnableContentEncoding(true)
	restful.TraceLogger(&httpLog{})
	restful.SetLogger(&httpLog{})
	return &HttpServer{
		Addr:      addr,
		Container: c,
		srv: &http.Server{
			Addr:    addr,
			Handler: c,
		},
	}
}

func (s *HttpServer) Type() Type {
	return TypeHTTP
}

func (s *HttpServer) Start(ctx context.Context) error {
	pkg.Log.Infof("HTTP Server listen: %s", s.Addr)
	go func() {
		if global.Conf.Server.Tls.Enable {
			pkg.Log.Infof("HTTPS Server listen: %s", s.Addr)
			if err := s.srv.ListenAndServeTLS(global.Conf.Server.Tls.CertFile, global.Conf.Server.Tls.KeyFile); err != nil {
				pkg.Log.Errorf("error http serve: %s", err)
			}
		} else {
			pkg.Log.Infof("HTTP Server listen: %s", s.Addr)
			if err := s.srv.ListenAndServe(); err != nil {
				pkg.Log.Errorf("error http serve: %s", err)
			}
		}
	}()
	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

type httpLog struct{}

func (t *httpLog) Print(v ...any) {
	pkg.Log.Debug(v...)
}

func (t *httpLog) Printf(format string, v ...any) {
	pkg.Log.Debugf(format, v...)
}

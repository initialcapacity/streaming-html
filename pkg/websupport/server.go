package websupport

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
)

type Server struct {
	server *http.Server
	port   int
}

type Handlers = func(mux *http.ServeMux)

func NewServer(handlers Handlers) *Server {
	mux := http.NewServeMux()

	handlers(mux)

	return &Server{
		server: &http.Server{Handler: mux},
	}
}

func (s *Server) Start(host string, port int) (listenerPort int, done chan error) {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		done <- err
		return 0, done
	}

	go func() {
		done <- s.server.Serve(listen)
	}()

	listenerPort = listen.Addr().(*net.TCPAddr).Port
	slog.Info("Server started", "address", fmt.Sprintf("http://%s:%d", host, listenerPort))
	s.port = listenerPort
	return listenerPort, done
}

func (s *Server) Stop() error {
	return s.server.Shutdown(context.Background())
}

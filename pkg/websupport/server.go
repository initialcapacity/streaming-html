package websupport

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"
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

func (s *Server) WaitUntilHealthy(path string) error {
	statusCode := make(chan int)

	go func() {
		for {
			resp, err := http.Get(fmt.Sprintf("http://localhost:%d%s", s.port, path))
			if err == nil {
				statusCode <- resp.StatusCode
				return
			}
		}
	}()

	select {
	case code := <-statusCode:
		if code == http.StatusOK {
			return nil
		} else {
			return fmt.Errorf("server responded with a non 200 code: %d", code)
		}
	case <-time.After(100 * time.Millisecond):
		return errors.New("server did not respond in 100 milliseconds")
	}
}

func (s *Server) Stop() error {
	return s.server.Shutdown(context.Background())
}

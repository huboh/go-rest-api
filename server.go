package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

const (
	defHost = ""
	defPort = "4000"
)

// Server Config
type ServerConfig struct {
	Port       string
	Host       string
	ReqHandler http.Handler
}

// NewServerConfig properly create a new Server Config instance and returns it
func NewServerConfig(h string, p string, rh http.Handler) *ServerConfig {
	if h == "" {
		h = defHost
	}

	if p == "" {
		p = defPort
	}

	if rh == nil {
		log.Fatal("no request handler specified")
	}

	return &ServerConfig{
		Port:       p,
		Host:       h,
		ReqHandler: rh,
	}
}

// Server
type Server struct {
	configs *ServerConfig
	httpSvr *http.Server
}

// NewServer initiates a new Server instance and returns it
func NewServer(configs *ServerConfig) *Server {
	if configs == nil {
		configs = NewServerConfig(
			os.Getenv("HOST"),
			os.Getenv("PORT"),
			http.DefaultServeMux,
		)
	}

	return &Server{
		configs: configs,
		httpSvr: &http.Server{
			Addr:    net.JoinHostPort(configs.Host, configs.Port),
			Handler: configs.ReqHandler,
		},
	}
}

func (s *Server) Stop() error {
	if (s != nil) && (s.httpSvr != nil) {
		return s.httpSvr.Close()
	}

	return nil
}

func (s *Server) Start() error {
	return s.httpSvr.ListenAndServe()
}

// Package server provides functionality for configuring and running an HTTP server.
package server

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

// Config holds the configuration settings for the HTTP server.
type Config struct {
	Port    string
	Host    string
	Addr    string
	Handler http.Handler
}

// NewConfig properly create a new Server Config instance and returns it
func NewConfig(h string, p string, rh http.Handler) *Config {
	if h == "" {
		h = defHost
	}

	if p == "" {
		p = defPort
	}

	if rh == nil {
		log.Fatal("no request handler specified")
	}

	return &Config{
		Port:    p,
		Host:    h,
		Addr:    net.JoinHostPort(h, p),
		Handler: rh,
	}
}

// Server represents an HTTP server, encapsulating the configuration and the underlying http.Server instance.
type Server struct {
	Configs *Config
	httpSvr *http.Server
}

// New initiates a new Server instance and returns it
func New(configs *Config) *Server {
	if configs == nil {
		configs = NewConfig(
			os.Getenv("HOST"),
			os.Getenv("PORT"),
			http.DefaultServeMux,
		)
	}

	return &Server{
		Configs: configs,
		httpSvr: &http.Server{
			Addr:    configs.Addr,
			Handler: configs.Handler,
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

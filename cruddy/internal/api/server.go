package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/knowit/gogo-gopher/cruddy/internal/logging"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

var (
	healthy int32
	ready   int32
)

type ArrayResponse []string
type MapResponse map[string]string

type Server struct {
	router  *mux.Router
	logger  *zap.Logger
	handler http.Handler
}

type Config struct {
	ServerShutdownTimeout time.Duration `mapstructure:"server-shutdown-timeout"`
}

func NewServer(logger *zap.Logger) (*Server, error) {
	s := &Server{
		router: mux.NewRouter(),
		logger: logger,
	}
	return s, nil
}

func (s *Server) registerHandlers() {
}

func (s *Server) registerMiddlewares() {
	loggingMiddleware := logging.NewLoggingMiddleware(s.logger)
	s.router.Use(loggingMiddleware.Handler)
}

func (s *Server) startServer() *http.Server {
	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  2 * 30 * time.Second,
		Handler:      s.handler,
	}

	go func() {
		s.logger.Info("Starting HTTP server", zap.String("addr", srv.Addr))
		err := srv.ListenAndServe()

		if err != http.ErrServerClosed {
			s.logger.Fatal("HTTP server crashed", zap.Error(err))
		}
	}()

	return srv
}

func (s *Server) startSecureServer() *http.Server {
	srv := &http.Server{
		Addr:         "localhost:8081",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  2 * 30 * time.Second,
		Handler:      s.handler,
	}

	//cert := "path-to-tls-cert"
	//key := "path-to-tls-key"

	go func() {
		s.logger.Info("Starting HTTPS server", zap.String("addr", srv.Addr))
		err := srv.ListenAndServe()

		if err != http.ErrServerClosed {
			s.logger.Fatal("HTTPS server crashed", zap.Error(err))
		}
	}()

	return srv
}

func (s *Server) printRoutes() {
	s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}

		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}

		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}

		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}

		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}

		fmt.Println()
		return nil
	})
}

func (s *Server) ListenAndServe() (*http.Server, *http.Server, *int32, *int32) {
	//ctx := context.Background() // has no tracer to init with context

	s.handler = s.router

	atomic.StoreInt32(&healthy, 1)
	atomic.StoreInt32(&ready, 1)

	s.registerHandlers()
	s.registerMiddlewares()

	s.printRoutes()

	// TODO: Load config here!

	srv := s.startServer()
	secureSrv := s.startSecureServer()

	// TODO: Tell k8s that we're ready!

	return srv, secureSrv, &healthy, &ready
}

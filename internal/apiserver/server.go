package apiserver

import (
	"errors"
	"log/slog"
	"net/http"

	middleware "github.com/MortalSC/FastGO/pkg/middleware"
	genericoptions "github.com/MortalSC/FastGO/pkg/options"
	"github.com/gin-gonic/gin"
)

type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
	Addr         string
}

type Server struct {
	cfg *Config
	srv *http.Server
}

func (cfg *Config) NewServer() (*Server, error) {
	// Create gin engine
	engine := gin.New()

	middlewares := []gin.HandlerFunc{
		gin.Recovery(),
		middleware.NoCache,
		middleware.Cors,
		middleware.RequestID(),
	}
	engine.Use(middlewares...)

	// register 404 handler
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "PageNotFound",
			"message": "Page not found",
		})
	})

	// register /healthz handler
	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// create HTTP server instance
	httpSrv := &http.Server{
		Addr:    cfg.Addr,
		Handler: engine,
	}

	return &Server{
		cfg: cfg,
		srv: httpSrv,
	}, nil
}

func (s *Server) Run() error {
	slog.Info("Read MySQL host from config", "mysql.addr", s.cfg.MySQLOptions.Addr)
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	select {}
}

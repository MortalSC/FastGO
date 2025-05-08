package apiserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	slog.Info("Start to listening the incoming requests on http address", "addr", s.cfg.Addr)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	// create a channel of type os.Signal for receiving system signals
	quit := make(chan os.Signal, 1)
	// When the kill command is executed (without parameters), the syscall.SIGTERM signal is sent by default
	// Using the kill-2 command will send the syscall.SIGINT signal (for example, triggered by pressing CTRL+C)
	// Using the kill-9 command will send the syscall.SIGKILL signal, but the SIGKILL signal cannot be captured, so there is no need to listen for and handle it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Block the program and wait to receive the signal from the quit channel
	<-quit

	slog.Info("Shutting down server...")

	// Gracefully close the service
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// First, close the dependent services, and then close the dependent services
	// Gracefully close the service within 10 seconds (process the unprocessed requests and then close the service). If it exceeds 10 seconds, it will time out and exit
	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error("Insecure Server forced to shutdown:", "err", err)
		return err
	}

	slog.Info("Server exited")

	return nil
}

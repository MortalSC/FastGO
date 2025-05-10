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

	"github.com/MortalSC/FastGO/internal/apiserver/biz"
	"github.com/MortalSC/FastGO/internal/apiserver/handler"
	"github.com/MortalSC/FastGO/internal/apiserver/store"
	"github.com/MortalSC/FastGO/internal/pkg/core"
	"github.com/MortalSC/FastGO/internal/pkg/errorx"
	"github.com/MortalSC/FastGO/internal/pkg/known"
	middleware "github.com/MortalSC/FastGO/internal/pkg/middleware"
	"github.com/MortalSC/FastGO/internal/pkg/validation"
	genericoptions "github.com/MortalSC/FastGO/pkg/options"
	"github.com/MortalSC/FastGO/pkg/token"
	"github.com/gin-gonic/gin"
)

type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
	Addr         string
	JWTKey       string
	ExpiraTime   time.Duration
}

type Server struct {
	cfg *Config
	srv *http.Server
}

func (cfg *Config) NewServer() (*Server, error) {
	token.Init(cfg.JWTKey, known.XUserID, cfg.ExpiraTime)

	// Create gin engine
	engine := gin.New()

	middlewares := []gin.HandlerFunc{
		gin.Recovery(),
		middleware.NoCache,
		middleware.Cors,
		middleware.RequestID(),
	}
	engine.Use(middlewares...)

	// Initialize database connection
	db, err := cfg.MySQLOptions.NewDB()
	if err != nil {
		return nil, err
	}
	store := store.NewStore(db)

	cfg.InstallRESTAPI(engine, store)

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

func (cfg *Config) InstallRESTAPI(engine *gin.Engine, store store.IStore) {

	// ====== test api start ======

	// register 404 handler
	engine.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errorx.ErrNotFound.WithMessage("Page not found"), nil)
	})

	// register /healthz handler
	engine.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, map[string]string{
			"status": "ok",
		}, nil)
	})

	// ====== test api end ======

	handler := handler.NewHandler(biz.NewBiz(store), validation.NewValidation(store))

	engine.POST("/login", handler.Login)
	engine.POST("/refresh-token", middleware.Authn(), handler.RefreshToken)

	authMiddleware := []gin.HandlerFunc{
		middleware.Authn(),
	}

	// Register the V1 API routes
	v1 := engine.Group("/api/v1")
	{
		userv1 := v1.Group("/user")
		{
			userv1.POST("", handler.CreateUser)
			userv1.Use(authMiddleware...)
			userv1.PUT(":user_id", handler.UpdateUser)
			userv1.DELETE(":user_id", handler.DeleteUser)
			userv1.GET(":user_id", handler.GetUser)
			userv1.GET("", handler.ListUsers)
			userv1.PUT(":user_id/change-password", handler.ChangePassword)
		}

		postv1 := v1.Group("/post", authMiddleware...)
		{
			postv1.POST("", handler.CreatePost)
			postv1.PUT(":post_id", handler.UpdatePost)
			postv1.DELETE(":post_id", handler.DeletePost)
			postv1.GET(":post_id", handler.GetPost)
			postv1.GET("", handler.ListPosts)
		}
	}
}

package apiserver

import (
	"log/slog"

	genericoptions "github.com/MortalSC/FastGO/pkg/options"
)

type Config struct {
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

type Server struct {
	cfg *Config
}

func (cfg *Config) NewServer() (*Server, error) {
	return &Server{
		cfg: cfg,
	}, nil
}

func (s *Server) Run() error {
	slog.Info("Read MySQL host from config", "mysql.addr", s.cfg.MySQLOptions.Addr)
	select {}
}

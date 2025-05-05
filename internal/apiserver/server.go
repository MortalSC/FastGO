package apiserver

import (
	"fmt"

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
	fmt.Printf("Read MySQL host from config: %s\n", s.cfg.MySQLOptions.Addr)
	select {}
}

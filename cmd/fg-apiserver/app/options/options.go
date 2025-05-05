package options

import (
	"github.com/MortalSC/FastGO/internal/apiserver"
	genericoptions "github.com/MortalSC/FastGO/pkg/options"
)

// ServerOptions contains all service configuration options
// Acts as a root configuration container that aggregates various subsystem configurations
// Supports both JSON serialization and configuration file parsing via mapstructure
type ServerOptions struct {
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions creates a ServerOptions instance with default values
// Initializes all subsystem configurations with their respective defaults
// Suitable for development environments. Production deployments should override
// these values through configuration files or environment variables
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: genericoptions.NewMySQLOptions(),
	}
}

// Validate performs full configuration validation
// Executes validation recursively for all subsystem configurations
// Returns the first encountered error or nil if all configurations are valid
// Ensures the service starts with a valid configuration state
func (s *ServerOptions) Validate() error {
	if err := s.MySQLOptions.Validate(); err != nil {
		return err
	}

	return nil
}

// Config converts ServerOptions to apiserver-ready configuration
// Transforms root configuration object into domain-specific configuration
// The returned Config object should be treated as immutable. Any modifications
// should be made through ServerOptions before regeneration
func (s *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: s.MySQLOptions,
	}, nil
}

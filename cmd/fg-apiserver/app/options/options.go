package options

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// MySQLOptions : defines options for mysql databsse
type MySQLOptions struct {
	Addr                  string        `json:"addr" mapstructure:"addr"`
	Username              string        `json:"username" mapstructure:"username"`
	Password              string        `json:"password" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections" mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time" mapstructure:"max-connection-life-time"`
}

// NewMySQLOptions : creates a new MysqlOptions object with default values
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:                  "127.0.0.1:3306",
		Username:              "xxx",
		Password:              "xxx",
		Database:              "xxx",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

type ServerOptions struct {
	MySQLOptions *MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions : creates a new ServerOptions object with default values
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: NewMySQLOptions(),
	}
}

// Validate : validates the ServerOptions object
func (s *ServerOptions) Validate() error {
	// validate MySQL options
	// check if the address is empty
	if s.MySQLOptions.Addr == "" {
		return fmt.Errorf("mysql addr is required")
	}
	// check if the host is invalid
	host, postStr, err := net.SplitHostPort(s.MySQLOptions.Addr)
	if err != nil {
		return fmt.Errorf("Invalid MySQL address format: '%s': %v", s.MySQLOptions.Addr, err)
	}
	// check if the port is invalid
	port, err := strconv.Atoi(postStr)
	if err != nil || port < 0 || port > 65535 {
		return fmt.Errorf("Invalid MySQL port: %s", postStr)
	}
	// check if the host is empty
	if host == "" {
		return fmt.Errorf("mysql host is required")
	}

	// check if the username is empty
	if s.MySQLOptions.Username == "" {
		return fmt.Errorf("mysql username is required")
	}
	// check if the password is empty
	if s.MySQLOptions.Password == "" {
		return fmt.Errorf("mysql password is required")
	}
	// check if the database is empty
	if s.MySQLOptions.Database == "" {
		return fmt.Errorf("mysql database is required")
	}
	// check if the max idle connections is less than 0
	if s.MySQLOptions.MaxIdleConnections <= 0 {
		return fmt.Errorf("mysql max idle connections must be greater than or equal to 0")
	}
	// check if the max open connections is less than 0
	if s.MySQLOptions.MaxOpenConnections <= 0 {
		return fmt.Errorf("mysql max open connections must be greater than or equal to 0")
	}
	// check if the max connection life time is less than 0
	if s.MySQLOptions.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("mysql max connection life time must be greater than or equal to 0")
	}

	return nil
}

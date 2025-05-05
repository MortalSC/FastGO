package options

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLOptions defines configuration options for MySQL database connections
// Fields are tagged for JSON serialization and configuration mapping using mapstructure
type MySQLOptions struct {
	Addr                  string        `json:"addr" mapstructure:"addr"`
	Username              string        `json:"username" mapstructure:"username"`
	Password              string        `json:"password" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections" mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time" mapstructure:"max-connection-life-time"`
}

// NewMySQLOptions creates a MySQLOptions instance with default values
// Defaults are suitable for development environment. Production deployments should override these values
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

// Validate checks the configuration options for validity
// Returns the first encountered error or nil if configuration is valid
func (s *MySQLOptions) Validate() error {
	// Validate address format
	if s.Addr == "" {
		return fmt.Errorf("mysql addr is required")
	}

	// Parse host and port components
	host, portStr, err := net.SplitHostPort(s.Addr)
	if err != nil {
		return fmt.Errorf("Invalid MySQL address format '%s': %v", s.Addr, err)
	}

	// Validate port number
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 0 || port > 65535 {
		return fmt.Errorf("Invalid MySQL port: %s", portStr)
	}

	// Validate host presence
	if host == "" {
		return fmt.Errorf("mysql host is required")
	}

	// Validate authentication credentials
	if s.Username == "" {
		return fmt.Errorf("mysql username is required")
	}
	if s.Password == "" {
		return fmt.Errorf("mysql password is required")
	}
	if s.Database == "" {
		return fmt.Errorf("mysql database is required")
	}

	// Validate connection pool parameters
	if s.MaxIdleConnections <= 0 {
		return fmt.Errorf("mysql max idle connections must be greater than or equal to 0")
	}
	if s.MaxOpenConnections <= 0 {
		return fmt.Errorf("mysql max open connections must be greater than or equal to 0")
	}
	if s.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("mysql max connection life time must be greater than or equal to 0")
	}

	return nil
}

// DSN constructs the MySQL Data Source Name string
// Format: username:password@protocol(address)/dbname?param=value
// Includes charset, time parsing, and timezone configuration
func (s *MySQLOptions) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		s.Username,
		s.Password,
		s.Addr,
		s.Database,
		true,
		"Local",
	)
}

// NewDB creates and configures a GORM database instance
// The returned *gorm.DB has connection pool parameters configured
// Caller is responsible for closing the database connection when done
func (s *MySQLOptions) NewDB() (*gorm.DB, error) {
	// Initialize GORM with MySQL driver
	db, err := gorm.Open(mysql.Open(s.DSN()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	// Get underlying *sql.DB instance for connection pool configuration
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configure connection pool settings
	sqlDB.SetMaxOpenConns(s.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(s.MaxIdleConnections)
	sqlDB.SetConnMaxLifetime(s.MaxConnectionLifeTime)

	return db, nil
}

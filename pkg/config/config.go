package config

import (
	"net/http"
	"time"

	"github.com/shurcooL/githubv4"
	"github.com/urfave/cli/v2"
)

// Server defines the Github API endpoint details
type Server struct {
	Addr    string
	Path    string
	Timeout time.Duration // timeout for prom /metrics endpoint
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string
	Pretty bool
}

// Config defines the application config
type Config struct {
	Client         *githubv4.Client
	RawClient      http.Client
	Server         Server
	Logs           Logs
	Token          string
	BaseURL        string
	Insecure       bool
	RequestTimeout time.Duration // request timeout calling Github API
	EnterpriseName cli.StringSlice
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}

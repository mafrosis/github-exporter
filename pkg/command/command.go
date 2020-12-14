package command

import (
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/urfave/cli/v2"
	"github.com/mafrosis/github-exporter/pkg/action"
	"github.com/mafrosis/github-exporter/pkg/config"
	"github.com/mafrosis/github-exporter/pkg/version"
)

var (
	// ErrMissingGithubToken defines the error if github.token is empty.
	ErrMissingGithubToken = `Missing required github.token`
)

// Run parses the command line arguments and executes the program.
func Run() error {
	cfg := config.Load()

	app := &cli.App{
		Name:    "ghe-exporter",
		Version: version.String,
		Usage:   "GitHub Enterprise Exporter",
		Authors: []*cli.Author{
			{
				Name:  "Matt Black",
				Email: "blackm@anz.com",
			},
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log.level",
				Value:       "info",
				Usage:       "Only log messages with given severity",
				EnvVars:     []string{"GITHUB_EXPORTER_LOG_LEVEL"},
				Destination: &cfg.Logs.Level,
			},
			&cli.BoolFlag{
				Name:        "log.pretty",
				Value:       false,
				Usage:       "Enable pretty messages for logging",
				EnvVars:     []string{"GITHUB_EXPORTER_LOG_PRETTY"},
				Destination: &cfg.Logs.Pretty,
			},
			&cli.StringFlag{
				Name:        "web.address",
				Value:       "0.0.0.0:9504",
				Usage:       "Address to bind the metrics server",
				EnvVars:     []string{"GITHUB_EXPORTER_WEB_ADDRESS"},
				Destination: &cfg.Server.Addr,
			},
			&cli.StringFlag{
				Name:        "web.path",
				Value:       "/metrics",
				Usage:       "Path to bind the metrics server",
				EnvVars:     []string{"GITHUB_EXPORTER_WEB_PATH"},
				Destination: &cfg.Server.Path,
			},
			&cli.DurationFlag{
				Name:        "web.timeout",
				Value:       10 * time.Second,
				Usage:       "Server metrics endpoint timeout",
				EnvVars:     []string{"GITHUB_EXPORTER_WEB_TIMEOUT"},
				Destination: &cfg.Server.Timeout,
			},
			&cli.DurationFlag{
				Name:        "request.timeout",
				Value:       5 * time.Second,
				Usage:       "Timeout requesting GitHub API",
				EnvVars:     []string{"GITHUB_EXPORTER_REQUEST_TIMEOUT"},
				Destination: &cfg.RequestTimeout,
			},
			&cli.StringFlag{
				Name:        "github.token",
				Value:       "",
				Usage:       "Access token for the GitHub API",
				EnvVars:     []string{"GITHUB_EXPORTER_TOKEN"},
				Destination: &cfg.Token,
			},
			&cli.StringFlag{
				Name:        "github.baseurl",
				Value:       "",
				Usage:       "URL to access the GitHub Enterprise API",
				EnvVars:     []string{"GITHUB_EXPORTER_BASE_URL"},
				Destination: &cfg.BaseURL,
			},
			&cli.BoolFlag{
				Name:        "github.insecure",
				Value:       false,
				Usage:       "Skip TLS verification for GitHub Enterprise",
				EnvVars:     []string{"GITHUB_EXPORTER_INSECURE"},
				Destination: &cfg.Insecure,
			},
			&cli.StringSliceFlag{
				Name:        "github.enterprise",
				Value:       cli.NewStringSlice(),
				Usage:       "Enterprise name to scrape metrics from",
				EnvVars:     []string{"GITHUB_EXPORTER_ENTERPRISE_NAME"},
				Destination: &cfg.EnterpriseName,
			},
		},
		Commands: []*cli.Command{
			Health(cfg),
		},
		Action: func(c *cli.Context) error {
			logger := setupLogger(cfg)

			if cfg.Token == "" {
				level.Error(logger).Log(
					"msg", ErrMissingGithubToken,
				)

				return fmt.Errorf(ErrMissingGithubToken)
			}

			return action.Server(cfg, logger)
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	return app.Run(os.Args)
}

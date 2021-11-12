package command

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log/level"
	"github.com/urfave/cli/v2"
	"github.com/mafrosis/github-exporter/pkg/config"
)

// Health provides the sub-command to perform a health check.
func Health(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "health",
		Usage: "Perform health checks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "web.address",
				Value:       "0.0.0.0:9504",
				Usage:       "Address to bind the metrics server",
				EnvVars:     []string{"GITHUB_EXPORTER_WEB_ADDRESS"},
				Destination: &cfg.Server.Addr,
			},
		},
		Action: func(c *cli.Context) error {
			logger := setupLogger(cfg)

			resp, err := http.Get(
				fmt.Sprintf(
					"http://%s/healthz",
					cfg.Server.Addr,
				),
			)

			if err != nil {
				level.Error(logger).Log(
					"msg", "Failed to request health check",
					"err", err,
				)

				return err
			}

			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				level.Error(logger).Log(
					"msg", "Health check seems to be in bad state",
					"err", err,
					"code", resp.StatusCode,
				)

				return err
			}

			return nil
		},
	}
}

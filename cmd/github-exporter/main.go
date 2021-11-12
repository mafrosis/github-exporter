package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mafrosis/github-exporter/pkg/command"
)

func main() {
	if env := os.Getenv("GITHUB_EXPORTER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}

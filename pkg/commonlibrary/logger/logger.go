package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/config"
)

func New(cfg *config.Config) *zap.Logger {
	var logger *zap.Logger

	var err error

	if cfg.Env == "prod" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		fmt.Printf("Failed to create logger: %v\n", err)
		os.Exit(1)
	}

	defer func() {
		_ = logger.Sync()
	}() // flushes any buffered log entries

	return logger
}

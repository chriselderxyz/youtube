package main

import (
	"fmt"
	"log/slog"
	"os"

	phuslog "github.com/phuslu/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/////////////////////////////////////////////////////////////
// Default Loggers
/////////////////////////////////////////////////////////////

func main() {

	// slog + phuslu/log
	fmt.Println("====================================================")
	fmt.Println("slog + phuslu/log handler")
	fmt.Println("====================================================")

	opt := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	ph := phuslog.SlogNewJSONHandler(os.Stdout, opt)
	spLogger := slog.New(ph)

	slog.SetDefault(spLogger)
	slog.Info("slog message")

	// zerolog
	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	zLogger = zLogger.With().
		Timestamp().
		Logger()

	log.Logger = zLogger

	log.Info().
		Msg("zerolog")
}

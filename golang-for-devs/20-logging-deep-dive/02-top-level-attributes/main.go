package main

import (
	"fmt"
	"log/slog"
	"os"

	phuslog "github.com/phuslu/log"
	"github.com/rs/zerolog"
)

/////////////////////////////////////////////////////////////
// Top Level Attributes
/////////////////////////////////////////////////////////////

func main() {

	fmt.Println("====================================================")
	fmt.Println("slog + phuslu/log handler")
	fmt.Println("====================================================")

	opt := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}
	ph := phuslog.SlogNewJSONHandler(os.Stdout, opt)
	spLogger := slog.New(ph)
	spLogger = spLogger.With("env", "dev", "id", "12345")
	spLogger.Info("slog + phuslu/log Handler")

	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Str("env", "dev").
		Str("id", "12345").
		Logger()

	zLogger.Info().Msg("zerolog")
}

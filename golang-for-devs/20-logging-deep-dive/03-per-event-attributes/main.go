package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	phuslog "github.com/phuslu/log"
	"github.com/rs/zerolog"
)

// ///////////////////////////////////////////////////////////
// Per Event Attributes
// ///////////////////////////////////////////////////////////
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
	spLogger.Info("slog + phuslu/log", "env", "dev", "num", 10)
	spLogger.LogAttrs(context.Background(), slog.LevelInfo, "slog message",
		slog.String("user", "12345"),
		slog.Int("num", 10),
		slog.Bool("false?", true),
	)

	// zerolog
	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	zLogger = zLogger.With().
		Timestamp().
		Logger()

	zLogger.Info().Str("user", "12345").Bool("is_valid", true).Msg("zerolog")
}

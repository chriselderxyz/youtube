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
// Nested JSON Objects
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

	spLogger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"slog + phuslu/log - with Attrs",
		slog.GroupAttrs("nested", slog.String("name", "chris"), slog.Int("numb", 10)),
	)

	// zerolog
	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	zLogger = zLogger.With().
		Timestamp().
		Logger()

	dict := zerolog.Dict().
		Str("name", "chris").
		Int("numb", 10)

	zLogger.Info().
		Dict("nested", dict).
		Msg("zerolog")
}

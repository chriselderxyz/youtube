package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	phuslog "github.com/phuslu/log"
	"github.com/rs/zerolog"
)

/////////////////////////////////////////////////////////////
// Hooks
/////////////////////////////////////////////////////////////

func main() {

	// slog + phuslu/log
	fmt.Println("====================================================")
	fmt.Println("slog + phuslu/log handler")
	fmt.Println("====================================================")

	opt := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			switch attr.Key {
			case slog.TimeKey:
				attr.Key = "timestamp"
			}

			return attr
		},
	}
	ph := phuslog.SlogNewJSONHandler(os.Stdout, opt)
	spLogger := slog.New(ph)

	spLogger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"slog + phuslu/log - with Attrs",
	)

	// zerolog
	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	zLogger = zLogger.With().
		Timestamp().
		Logger()

	hook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		if level >= zerolog.ErrorLevel {
			e.Bool("alert", true)
		}
	})

	zLogger = zLogger.Hook(hook)

	zLogger.Error().
		Msg("zerolog")
}

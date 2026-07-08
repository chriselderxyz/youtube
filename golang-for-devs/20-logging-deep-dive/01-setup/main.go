package main

import (
	"fmt"
	"log/slog"
	"os"

	phuslog "github.com/phuslu/log"
	"github.com/rs/zerolog"
)

/////////////////////////////////////////////////////////////
// Basic Logging
/////////////////////////////////////////////////////////////

// slog levels
// const (
// 	LelDebug Lveevel = -4
// 	LevelInfo  Level = 0
// 	LevelWarn  Level = 4
// 	LevelError Level = 8
// )

func main() {
	fmt.Println("====================================================")
	fmt.Println("slog")
	fmt.Println("====================================================")
	
	opt := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	h := slog.NewJSONHandler(os.Stdout, opt)

	sLogger := slog.New(h)

	sLogger.Info("message from slog")

	fmt.Println("====================================================")
	fmt.Println("slog + phuslu/log - Handler")
	fmt.Println("====================================================")
	
	opt2 := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	h2 := phuslog.SlogNewJSONHandler(os.Stdout, opt2)

	spLogger := slog.New(h2)

	spLogger.Warn("slog + phuslu/log Handler")
	
	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	zLogger.Info().Msg("zerolog message")
}

package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	phuslog "github.com/phuslu/log"
	"github.com/rs/zerolog"
)

/////////////////////////////////////////////////////////////
// Lists
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

	spLogger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"slog + phuslu/log - with Attrs",
		slog.Any("list", []int{1,2,3,4}),
	)

	// zerolog
	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	zLogger = zLogger.With().
		Timestamp().
		Logger()

	arr := zerolog.Arr().Str("name").Bool(true).Int(10)

	zLogger.Info().
		Strs("strings", []string{"1", "2", "3"}).
		Durs("duration", []time.Duration{time.Second, time.Microsecond}).
		Array("list", arr).
		Msg("zerolog")
}

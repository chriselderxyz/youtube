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

//////////////////////////////////////////////////////////////
// Sampling
//////////////////////////////////////////////////////////////

type MyHandler struct {
	next  slog.Handler
	count int
}

func (h *MyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.next.Enabled(ctx, level)
}

func (h *MyHandler) Handle(ctx context.Context, r slog.Record) error {
	if h.count%2 == 0 {
		h.count++
		return h.next.Handle(ctx, r)
	}

	h.count++
	fmt.Println("Log Skipped")
	return nil
}

func (h *MyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &MyHandler{
		next:  h.next.WithAttrs(attrs),
		count: h.count,
	}
}

func (h *MyHandler) WithGroup(name string) slog.Handler {
	return &MyHandler{
		next:  h.next.WithGroup(name),
		count: h.count,
	}
}

func main() {

	// slog + phuslu/log
	fmt.Println("====================================================")
	fmt.Println("slog + phuslu/log handler")
	fmt.Println("====================================================")

	opt := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	ph := phuslog.SlogNewJSONHandler(os.Stdout, opt)

	wrapper := &MyHandler{
		next:  ph,
		count: 0,
	}

	spLogger := slog.New(wrapper)

	slog.SetDefault(spLogger)
	slog.Info("slog message")
	slog.Info("Sampled Out")

	// zerolog
	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	zLogger = zLogger.With().
		Timestamp().
		Logger()

	zLogger = zLogger.Sample(zerolog.LevelSampler{
		InfoSampler:  &zerolog.BasicSampler{N: 10},
		DebugSampler: zerolog.RandomSampler(100),
		WarnSampler: &zerolog.BurstSampler{
			Burst:  5,
			Period: time.Second,
			NextSampler: &zerolog.Sometimes,
		},
	})

	zLogger.Info().
		Msg("zerolog")
}

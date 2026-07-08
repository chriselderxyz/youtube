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
// Logging Custom Types
/////////////////////////////////////////////////////////////

type User struct {
	id       string
	age_days int
}

func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", u.id),
		slog.Int("age_days", u.age_days),
	)
}

func (u User) MarshalZerologObject(e *zerolog.Event) {
	e.Str("id", u.id).Int("age_days", u.age_days)
}

func main() {

	user := User{
		id:       "12345",
		age_days: 200,
	}

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
		slog.Any("user", user),
	)

	// zerolog
	fmt.Println("====================================================")
	fmt.Println("zerolog")
	fmt.Println("====================================================")

	zLogger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	zLogger = zLogger.With().
		Timestamp().
		Logger()

	zLogger.Info().
		Object("user", user).
		Msg("zerolog")

	zLogger.Info().
		EmbedObject(user).
		Msg("zerolog - embed")
}

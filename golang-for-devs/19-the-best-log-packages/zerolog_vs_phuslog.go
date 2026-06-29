package main

import (
	"runtime"
	"time"

	phuslulog "github.com/phuslu/log"
	"github.com/rs/zerolog"
	zerologlog "github.com/rs/zerolog/log"
)

/////////////////////////////////////////////////////////////
// zerolog vs phuslu/log
/////////////////////////////////////////////////////////////

func main() {

	// phuslu/log - Easy Multi-Level Writer

	phuslulog.DefaultLogger.Writer = &phuslulog.MultiLevelWriter{
		InfoWriter:    &phuslulog.FileWriter{Filename: "main.INFO", MaxSize: 100 << 20},
		WarnWriter:    &phuslulog.FileWriter{Filename: "main.WARNING", MaxSize: 100 << 20},
		ErrorWriter:   &phuslulog.FileWriter{Filename: "main.ERROR", MaxSize: 100 << 20},
		ConsoleWriter: &phuslulog.ConsoleWriter{ColorOutput: true},
		ConsoleLevel:  phuslulog.ErrorLevel,
	}

	phuslulog.Info().Int("number", 42).Str("foo", "bar").Msg("a info log")
	phuslulog.Warn().Int("number", 42).Str("foo", "bar").Msg("a warn log")
	phuslulog.Error().Int("number", 42).Str("foo", "bar").Msg("a error log")

	// phuslu/log - Rotating File Writer

	logger := phuslulog.Logger{
		Level: phuslulog.ParseLevel("info"),
		Writer: &phuslulog.FileWriter{
			Filename:     "logs/main.log",
			FileMode:     0600,
			MaxSize:      100 * 1024 * 1024,
			MaxBackups:   7,
			EnsureFolder: true,
			LocalTime:    true,
		},
	}

	logger.Info().Msg("hello world")

	// zerolog - Sampling

	// Will let 5 debug messages per period of 1 second.
	// Over 5 debug message, 1 every 100 debug messages are logged.
	// Other levels are not sampled.
	sampled := zerologlog.Sample(zerolog.LevelSampler{
		DebugSampler: &zerolog.BurstSampler{
			Burst:       5,
			Period:      1 * time.Second,
			NextSampler: &zerolog.BasicSampler{N: 100},
		},
	})
	sampled.Debug().Msg("hello world")

	// zerolog - hooks
	hooked := zerologlog.Hook(SeverityHook{})
	hooked.Warn().Msg("")

}

// zerolog - Hooks
type SeverityHook struct{}

func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {

	// Add "severity" attribute
	if level != zerolog.NoLevel {
		e.Str("severity", level.String())
	}

	// Add field for errors and up only
	if level >= zerolog.ErrorLevel {
		e.Str("alert_group", "backend")
	}

	// Add runtime stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	e.Uint64("heap_bytes", m.HeapAlloc)
}

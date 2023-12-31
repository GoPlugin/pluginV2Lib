package logging

import (
	"os"
	"sync"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/config"
)

// People often call this multiple times
var loggingMu sync.Mutex

func getTestLevel() zerolog.Level {
	lvlStr := os.Getenv(config.EnvVarLogLevel)
	if lvlStr == "" {
		lvlStr = "info"
	}
	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		panic(err)
	}
	return lvl
}

func Init() {
	loggingMu.Lock()
	defer loggingMu.Unlock()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger().Level(getTestLevel())
}

func GetTestLogger(t *testing.T) zerolog.Logger {
	return zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger().Level(getTestLevel())
}

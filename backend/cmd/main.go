package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/mateuszGorczany/BESTVoteliator/internal/http/rest"
	common "github.com/mateuszGorczany/BESTVoteliator/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ErrorLogInit = errors.New("cannot initalize logger")
)

var logger *zap.Logger

func initLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := loggerConfig.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Errorf("%v, %v", ErrorLogInit, err).Error())
	}
	return logger
}

func main() {
	common.LoadConfig()
	common.Logger = initLogger()
	rest.Run()
}

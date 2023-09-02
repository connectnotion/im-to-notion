package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/api7/gopkg/pkg/log"
	dingtalklogger "github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"

	"github.com/ronething/im-to-notion/internal/pkg/config"
)

// Dief generates the error message according to the template and args, outputs them to
// stderr and then exit the program with the code 1.
func Dief(template string, args ...interface{}) {
	if !strings.HasSuffix(template, "\n") {
		template += "\n"
	}
	fmt.Fprintf(os.Stderr, template, args...)
	os.Exit(1)
}

// SetDefaultLogger changes the log.DefaultLogger object with the given log configs.
func SetDefaultLogger(cfg config.LogOptions) error {
	logger, err := log.NewLogger(
		log.WithLogLevel(cfg.Level),
		log.WithOutputFile(cfg.Output),
		log.WithSkipFrames(4),
	)
	if err != nil {
		return err
	}
	log.DefaultLogger = logger
	dl := Logger{logger}
	dingtalklogger.SetLogger(&dl)
	return nil
}

type Logger struct {
	*log.Logger
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args)
	l.Warn(msg)
}

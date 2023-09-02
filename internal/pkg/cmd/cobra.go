package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/api7/gopkg/pkg/log"

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
	return nil
}

package main

import (
	"github.com/api7/gopkg/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	internalcmd "github.com/ronething/im-to-notion/internal/pkg/cmd"
	"github.com/ronething/im-to-notion/internal/pkg/config"
	"github.com/ronething/im-to-notion/internal/version"
)

const (
	_help = `im-to-notion is a tool that send msg to notion through im. You can
run by specifying the -c or --config option to specifying
configurations for this tool.
	`
)

var (
	_configFile string
)

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "im-to-notion [flags]",
		Long:    _help,
		Version: version.V.String(),
		Run: func(cmd *cobra.Command, args []string) {
			config.ViperDefaultValue()
			viper.SetConfigFile(_configFile)
			if err := viper.ReadInConfig(); err != nil {
				internalcmd.Dief(err.Error())
			}

			logOptions := config.LogOptions{}
			if err := viper.UnmarshalKey("log", &logOptions); err != nil {
				internalcmd.Dief(err.Error())
			}
			if err := internalcmd.SetDefaultLogger(logOptions); err != nil {
				internalcmd.Dief(err.Error())
			}

			log.Infow("config is",
				zap.Any("log", viper.Get("log")),
				zap.Any("source", viper.Get("source")),
				zap.Any("destination", viper.Get("destination")),
			)

		},
	}

	cmd.PersistentFlags().StringVarP(&_configFile, "config", "c", "./conf/config.yaml", "configuration file")
	return cmd
}

func main() {
	cmd := newCommand()
	if err := cmd.Execute(); err != nil {
		internalcmd.Dief(err.Error())
	}
}
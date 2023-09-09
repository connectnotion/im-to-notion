package config

import (
	"os"

	"github.com/spf13/viper"
)

// ViperDefaultValue set default config
func ViperDefaultValue() {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.output", "stderr")
}

// SetValueFromEnv set config from env
func SetValueFromEnv() error {
	var (
		logOptions  LogOptions
		source      Source
		destination Destination
	)

	if err := viper.UnmarshalKey("log", &logOptions); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("source", &source); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("destination", &destination); err != nil {
		return err
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if len(logLevel) > 0 {
		logOptions.Level = logLevel
	}
	logOutput := os.Getenv("LOG_OUTPUT")
	if len(logOutput) > 0 {
		logOptions.Output = logOutput
	}
	viper.Set("log", logOptions)

	dingtalkAppKey := os.Getenv("DINGTALK_APP_KEY")
	if len(dingtalkAppKey) > 0 {
		if source.Dingtalk == nil {
			source.Dingtalk = &Dingtalk{
				AppKey: dingtalkAppKey,
			}
		} else {
			source.Dingtalk.AppKey = dingtalkAppKey
		}
	}
	dingtalkAppSecret := os.Getenv("DINGTALK_APP_SECRET")
	if len(dingtalkAppSecret) > 0 {
		source.Dingtalk.AppSecret = dingtalkAppSecret
	}
	viper.Set("source", source)

	notionSecret := os.Getenv("NOTION_SECRET")
	if len(notionSecret) > 0 {
		destination.Notion.Secret = notionSecret
	}
	notionDatabaseId := os.Getenv("NOTION_DATABASE_ID")
	if len(notionDatabaseId) > 0 {
		destination.Notion.DatabaseId = notionDatabaseId
	}
	viper.Set("destination", destination)

	return nil
}

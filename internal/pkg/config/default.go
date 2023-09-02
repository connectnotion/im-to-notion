package config

import "github.com/spf13/viper"

func ViperDefaultValue() {

	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.output", "stderr")

	viper.SetDefault("source.dingtalk.appKey", "")
	viper.SetDefault("source.dingtalk.appSecret", "")

	viper.SetDefault("destination.notion.secret", "")
}

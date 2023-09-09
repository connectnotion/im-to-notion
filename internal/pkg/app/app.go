package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	internalcmd "github.com/ronething/im-to-notion/internal/pkg/cmd"
	"github.com/ronething/im-to-notion/internal/pkg/config"
	"github.com/ronething/im-to-notion/pkg/dingtalk"
	"github.com/ronething/im-to-notion/pkg/notion"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	// TODO: move to APP fields
	var (
		source      config.Source
		destination config.Destination
	)
	if err := viper.UnmarshalKey("source", &source); err != nil {
		internalcmd.Dief(err.Error())
	}
	if !source.Validate() {
		internalcmd.Dief("source config validate err, please check config")
	}
	if err := viper.UnmarshalKey("destination", &destination); err != nil {
		internalcmd.Dief(err.Error())
	}
	if !destination.Validate() {
		internalcmd.Dief("destination config validate err, please check config")
	}

	n := notion.NewNotion(destination.Notion.Secret, destination.Notion.DatabaseId)
	d := dingtalk.NewDingtalk(source.Dingtalk.AppKey, source.Dingtalk.AppSecret)
	d.SetNotion(n)
	// need to register function first
	d.RegisterFunction()

	if err := d.Start(context.TODO()); err != nil {
		internalcmd.Dief(err.Error())
	}
}

func (a *App) WaitForSignal() {
	stopC := make(chan os.Signal, 1)
	signal.Notify(stopC, syscall.SIGINT, syscall.SIGTERM)
	<-stopC
}

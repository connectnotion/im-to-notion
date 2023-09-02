package dingtalk

import (
	"context"
	"fmt"
	"strings"

	"github.com/api7/gopkg/pkg/log"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/payload"
	"go.uber.org/zap"

	"github.com/ronething/im-to-notion/pkg/notion"
)

type Dingtalk struct {
	Notion *notion.Notion
	Client *client.StreamClient
}

func NewDingtalk(appKey, appSecret string) *Dingtalk {
	cli := client.NewStreamClient(client.WithAppCredential(client.NewAppCredentialConfig(appKey, appSecret)))
	return &Dingtalk{Client: cli}
}

func (d *Dingtalk) SetNotion(n *notion.Notion) {
	d.Notion = n
}

// RegisterFunction register function for received event and callback
func (d *Dingtalk) RegisterFunction() {
	d.Client.RegisterAllEventRouter(d.OnEventReceived)
	d.Client.RegisterChatBotCallbackRouter(d.OnChatBotMessageReceived)
}

func (d *Dingtalk) Start(ctx context.Context) error {
	return d.Client.Start(ctx)
}

func (d *Dingtalk) Close() {
	d.Client.Close()
}

// OnEventReceived received header
func (*Dingtalk) OnEventReceived(ctx context.Context, df *payload.DataFrame) (frameResp *payload.DataFrameResponse, err error) {
	eventHeader := event.NewEventHeaderFromDataFrame(df)
	log.Infow("received event", zap.Any("eventHeader", eventHeader), zap.String("data", df.Data))

	frameResp = payload.NewSuccessDataFrameResponse()
	if err = frameResp.SetJson(event.NewEventProcessResultSuccess()); err != nil {
		return nil, err
	}

	return frameResp, nil
}

func (d *Dingtalk) OnChatBotMessageReceived(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
	var (
		msg     string
		title   string
		url     string
		comment string
	)
	log.Debugw("msg is", zap.Any("data", data))
	content := strings.TrimSpace(data.Text.Content)

	chatbotReplier := chatbot.NewChatbotReplier()
	// send to notion
	fields := strings.Split(content, "\n")
	// if content is empty, strings.Split(content, "\n") will also have len -> 1
	if len(fields) == 0 || len(content) == 0 {
		// nothing
		msg = "nothing send to notion"
		if err := chatbotReplier.SimpleReplyText(ctx, data.SessionWebhook, []byte(msg)); err != nil {
			return nil, err
		}
		return nil, nil
	}

	if len(fields) == 1 {
		if strings.HasPrefix(fields[0], "http") {
			url = fields[0]
		} else {
			title = fields[0]
		}
	}
	if len(fields) == 2 {
		if strings.HasPrefix(fields[0], "http") {
			url = fields[0]
			title = fields[1]
		} else {
			title = fields[0]
			url = fields[1]
		}
	}
	if len(fields) >= 3 {
		if strings.HasPrefix(fields[0], "http") {
			url = fields[0]
			title = fields[1]
		} else {
			title = fields[0]
			url = fields[1]
		}
		for i := 2; i < len(fields); i++ {
			comment += fields[i] + "\n"
		}
	}
	page, err := d.Notion.CreatePage(title, url, comment)
	if err != nil {
		log.Errorw("create page", zap.Any("err", err))
		return nil, err
	}
	log.Debugw("page", zap.Any("notion page", page))
	msg = fmt.Sprintf("send to notion, page url: %s", page.URL)
	if err := chatbotReplier.SimpleReplyText(ctx, data.SessionWebhook, []byte(msg)); err != nil {
		return nil, err
	}

	return nil, nil
}

package notion

import (
	"context"
	"time"

	"github.com/jomei/notionapi"
	"github.com/pkg/errors"
)

type Notion struct {
	//Client is notion api client
	Client *notionapi.Client
	//DatabaseId is notion database id which need to create page
	DatabaseId string
}

func NewNotion(token string, databaseId string) *Notion {
	return &Notion{
		Client:     notionapi.NewClient(notionapi.Token(token)),
		DatabaseId: databaseId,
	}
}

// CreatePage crate notion page
func (n *Notion) CreatePage(title, url, comment string) (*notionapi.Page, error) {
	_, err := n.Client.Database.Get(context.TODO(), notionapi.DatabaseID(n.DatabaseId))
	if err != nil {
		return nil, errors.Wrap(err, "get notion database")
	}

	now := notionapi.Date(time.Now())
	createReq := &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       "database_id",
			DatabaseID: notionapi.DatabaseID(n.DatabaseId),
		},
		Properties: notionapi.Properties{
			"Date": notionapi.DateProperty{
				Date: &notionapi.DateObject{
					Start: &now,
				},
			},
		},
	}
	if len(title) > 0 {
		createReq.Properties["Name"] = notionapi.TitleProperty{
			Title: []notionapi.RichText{
				{
					Text: &notionapi.Text{
						Content: title,
					},
				},
			},
		}
	}
	if len(url) > 0 {
		createReq.Properties["URL"] = notionapi.URLProperty{
			URL: url,
		}
	}
	if len(comment) > 0 {
		createReq.Properties["Comment"] = notionapi.RichTextProperty{
			RichText: []notionapi.RichText{
				{
					Text: &notionapi.Text{
						Content: comment,
					},
				},
			},
		}
	}
	newPage, err := n.Client.Page.Create(context.TODO(), createReq)
	if err != nil {
		return nil, errors.Wrap(err, "create notion page")
	}

	return newPage, nil
}

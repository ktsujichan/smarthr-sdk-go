package smarthr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Webhook struct {
	ID           string `json:"id" validate:"omitempty,uuid"`
	URL          string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	SecretToken  string `json:"secret_token,omitempty"`
	CrewCreated  bool   `json:"crew_created"`
	CrewUpdated  bool   `json:"crew_updated"`
	CrewDeleted  bool   `json:"crew_deleted"`
	CrewImported bool   `json:"crew_imported"`
	DisabledAt   string `json:"disabled_at,omitempty"`
	UpdatedAt    string `json:"updated_at"`
	CreatedAt    string `json:"created_at"`
}

func (c *Client) DeleteWebhook(ctx context.Context, id string) error {
	p := fmt.Sprintf("/v1/webhooks/%s", id)
	_, err := c.delete(ctx, p)
	return errors.WithStack(err)
}

func (c *Client) GetWebhook(ctx context.Context, id string) (*Webhook, error) {
	p := fmt.Sprintf("/v1/webhooks/%s", id)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var webhook Webhook
	if err := decodeBody(res, &webhook); err != nil {
		return nil, errors.WithStack(err)
	}
	return &webhook, nil
}

func (c *Client) UpdateWebhook(ctx context.Context, id string, data map[string]interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.patch(ctx, fmt.Sprintf("/v1/webhooks/%s", id), bytes.NewBuffer(b))
	return errors.WithStack(err)
}

type ListWebhooksOptions struct {
	Page    uint `url:"page,omitempty"`
	PerPage uint `url:"per_page,omitempty"`
}

func (opt *ListWebhooksOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListWebhooks(ctx context.Context, opt *ListWebhooksOptions) ([]Webhook, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/webhooks", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var webhooks []Webhook
	if err := decodeBody(res, &webhooks); err != nil {
		return nil, errors.WithStack(err)
	}
	return webhooks, nil
}

func (c *Client) CreateWebhook(ctx context.Context, webhook *Webhook) error {
	if webhook == nil {
		return nil
	}
	if err := validator.Struct(webhook); err != nil {
		return errors.WithStack(err)
	}
	b, err := json.Marshal(webhook)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.post(ctx, "/v1/webhooks", bytes.NewBuffer(b))
	return errors.WithStack(err)
}

package smarthr

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type MailFormat struct {
	ID             string          `json:"id"`
	MailType       string          `json:"mail_type"`
	Name           string          `json:"name"`
	CrewInputForms []CrewInputForm `json:"crew_input_forms,omitempty"`
	UpdatedAt      string          `json:"updated_at"`
	CreatedAt      string          `json:"created_at"`
}

func (c *Client) GetMailFormat(ctx context.Context, id string) (*MailFormat, error) {
	p := fmt.Sprintf("/v1/mail_formats/%s", id)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var mailFormat MailFormat
	if err := decodeBody(res, &mailFormat); err != nil {
		return nil, errors.WithStack(err)
	}
	return &mailFormat, nil
}

type ListMailFormatsOptions struct {
	Page    uint   `url:"page,omitempty"`
	PerPage uint   `url:"per_page,omitempty"`
	Embed   string `url:"embed,omitempty" validate:"omitempty,eq=crew_input_forms"`
}

func (opt *ListMailFormatsOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListMailFormats(ctx context.Context, opt *ListMailFormatsOptions) ([]MailFormat, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/mail_formats", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var mailFormats []MailFormat
	if err := decodeBody(res, &mailFormats); err != nil {
		return nil, errors.WithStack(err)
	}
	return mailFormats, nil
}

package smarthr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type CrewCustomFieldTemplate struct {
	ID        string                               `json:"id" validate:"omitempty,uuid"`
	Name      string                               `json:"name"`
	Type      string                               `json:"type" validate:"omitempty,crew_custom_field_template_type"`
	Elements  []CrewCustomEnumFieldTemplateElement `json:"elements,omitempty"`
	GroupID   string                               `json:"group_id" validate:"omitempty,uuid"`
	Group     *CrewCustomFieldTemplateGroup        `json:"group,omitempty"`
	Hint      string                               `json:"hint,omitempty"`
	Scale     int                                  `json:"scale,omitempty"`
	Position  int                                  `json:"position,omitempty"`
	UpdatedAt string                               `json:"updated_at"`
	CreatedAt string                               `json:"created_at"`
}

func (c *Client) DeleteCrewCustomFieldTemplate(ctx context.Context, id string) error {
	p := fmt.Sprintf("/v1/crew_custom_field_templates/%s", id)
	_, err := c.delete(ctx, p)
	return errors.WithStack(err)
}

func (c *Client) GetCrewCustomFieldTemplate(ctx context.Context, id string) (*CrewCustomFieldTemplate, error) {
	p := fmt.Sprintf("/v1/crew_custom_field_templates/%s", id)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var template CrewCustomFieldTemplate
	if err := decodeBody(res, &template); err != nil {
		return nil, errors.WithStack(err)
	}
	return &template, nil
}

func (c *Client) UpdateCrewCustomFieldTemplate(ctx context.Context, id string, data map[string]interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.patch(ctx, fmt.Sprintf("/v1/crew_custom_field_templates/%s", id), bytes.NewBuffer(b))
	return errors.WithStack(err)
}

type ListCrewCustomFieldTemplatesOptions struct {
	Page    uint   `url:"page,omitempty"`
	PerPage uint   `url:"per_page,omitempty"`
	Embed   string `url:"embed,omitempty" validate:"omitempty,eq=groups"`
}

func (opt *ListCrewCustomFieldTemplatesOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListCrewCustomFieldTemplates(ctx context.Context, opt *ListCrewCustomFieldTemplatesOptions) ([]CrewCustomFieldTemplate, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/crew_custom_field_templates", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var templates []CrewCustomFieldTemplate
	if err := decodeBody(res, &templates); err != nil {
		return nil, errors.WithStack(err)
	}
	return templates, nil
}

func (c *Client) CreateCrewCustomFieldTemplate(ctx context.Context, template *CrewCustomFieldTemplate) error {
	if template == nil {
		return nil
	}
	if err := validator.Struct(template); err != nil {
		return errors.WithStack(err)
	}
	b, err := json.Marshal(template)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.post(ctx, "/v1/crew_custom_field_templates", bytes.NewBuffer(b))
	return errors.WithStack(err)
}

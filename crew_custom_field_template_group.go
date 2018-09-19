package smarthr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type CrewCustomFieldTemplateGroup struct {
	ID         string                   `json:"id,omitempty"`
	Name       string                   `json:"name"`
	Position   int                      `json:"position,omitempty"`
	AccessType string                   `json:"access_type,omitempty" validate:"omitempty,eq=read_and_update|eq=hidden|eq=read_and_update_values"`
	Templates  *CrewCustomFieldTemplate `json:"templates,omitempty"`
	UpdatedAt  string                   `json:"updated_at"`
	CreatedAt  string                   `json:"created_at"`
}

func (c *Client) DeleteCrewCustomFieldTemplateGroup(ctx context.Context, id string) error {
	p := fmt.Sprintf("/v1/crew_custom_field_template_groups/%s", id)
	_, err := c.delete(ctx, p)
	return errors.WithStack(err)
}

func (c *Client) GetCrewCustomFieldTemplateGroup(ctx context.Context, id string) (*CrewCustomFieldTemplateGroup, error) {
	p := fmt.Sprintf("/v1/crew_custom_field_template_groups/%s", id)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var group CrewCustomFieldTemplateGroup
	if err := decodeBody(res, &group); err != nil {
		return nil, errors.WithStack(err)
	}
	return &group, nil
}

func (c *Client) UpdateCrewCustomFieldTemplateGroup(ctx context.Context, id string, data map[string]interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.patch(ctx, fmt.Sprintf("/v1/crew_custom_field_template_groups/%s", id), bytes.NewBuffer(b))
	return errors.WithStack(err)
}

type ListCrewCustomFieldTemplateGroupsOptions struct {
	Page    uint   `url:"page,omitempty"`
	PerPage uint   `url:"per_page,omitempty"`
	Embed   string `url:"embed,omitempty" validate:"omitempty,eq=templates"`
}

func (opt *ListCrewCustomFieldTemplateGroupsOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListCrewCustomFieldTemplateGroups(ctx context.Context, opt *ListCrewCustomFieldTemplateGroupsOptions) ([]CrewCustomFieldTemplateGroup, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/crew_custom_field_template_groups", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var groups []CrewCustomFieldTemplateGroup
	if err := decodeBody(res, &groups); err != nil {
		return nil, errors.WithStack(err)
	}
	return groups, nil
}

func (c *Client) CreateCrewCustomFieldTemplateGroup(ctx context.Context, group *CrewCustomFieldTemplateGroup) error {
	if group == nil {
		return nil
	}
	if err := validator.Struct(group); err != nil {
		return errors.WithStack(err)
	}
	b, err := json.Marshal(group)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.post(ctx, "/v1/crew_custom_field_template_groups", bytes.NewBuffer(b))
	return errors.WithStack(err)
}

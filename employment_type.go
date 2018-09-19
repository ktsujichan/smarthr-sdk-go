package smarthr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type EmploymentType struct {
	ID         string `json:"id" validate:"omitempty,uuid"`
	Name       string `json:"name"`
	PresetType string `json:"preset_type,omitempty" validate:"omitempty,emp_type"`
	UpdatedAt  string `json:"updated_at"`
	CreatedAt  string `json:"created_at"`
}

func (c *Client) DeleteEmploymentType(ctx context.Context, id string) error {
	p := fmt.Sprintf("/v1/employment_types/%s", id)
	_, err := c.delete(ctx, p)
	return errors.WithStack(err)
}

func (c *Client) GetEmploymentType(ctx context.Context, id string) (*EmploymentType, error) {
	p := fmt.Sprintf("/v1/employment_types/%s", id)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var employmentType EmploymentType
	if err := decodeBody(res, &employmentType); err != nil {
		return nil, errors.WithStack(err)
	}
	return &employmentType, nil
}

func (c *Client) UpdateEmploymentType(ctx context.Context, id string, data map[string]interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.patch(ctx, fmt.Sprintf("/v1/employment_types/%s", id), bytes.NewBuffer(b))
	return errors.WithStack(err)
}

type ListEmploymentTypesOptions struct {
	Page    uint `url:"page,omitempty"`
	PerPage uint `url:"per_page,omitempty"`
}

func (opt *ListEmploymentTypesOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListEmploymentTypes(ctx context.Context, opt *ListEmploymentTypesOptions) ([]EmploymentType, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/employment_types", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var employmentTypes []EmploymentType
	if err := decodeBody(res, &employmentTypes); err != nil {
		return nil, errors.WithStack(err)
	}
	return employmentTypes, nil
}

func (c *Client) CreateEmploymentType(ctx context.Context, employmentType *EmploymentType) error {
	if employmentType == nil {
		return nil
	}
	if err := validator.Struct(employmentType); err != nil {
		return errors.WithStack(err)
	}
	b, err := json.Marshal(employmentType)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.post(ctx, "/v1/employment_types", bytes.NewBuffer(b))
	return errors.WithStack(err)
}

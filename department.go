package smarthr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Department struct {
	ID        string       `json:"id" validate:"omitempty,uuid"`
	Name      string       `json:"name"`
	Position  int          `json:"position"`
	Code      string       `json:"code,omitempty"`
	Parent    *Department  `json:"parent,omitempty"`
	Children  []Department `json:"children,omitempty"`
	UpdatedAt string       `json:"updated_at"`
	CreatedAt string       `json:"created_at"`
}

func (c *Client) DeleteDepartment(ctx context.Context, id string) error {
	p := fmt.Sprintf("/v1/departments/%s", id)
	_, err := c.delete(ctx, p)
	return errors.WithStack(err)
}

func (c *Client) GetDepartment(ctx context.Context, id string) (*Department, error) {
	p := fmt.Sprintf("/v1/departments/%s", id)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var department Department
	if err := decodeBody(res, &department); err != nil {
		return nil, errors.WithStack(err)
	}
	return &department, nil
}

func (c *Client) UpdateDepartment(ctx context.Context, id string, data map[string]interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.patch(ctx, fmt.Sprintf("/v1/departments/%s", id), bytes.NewBuffer(b))
	return errors.WithStack(err)
}

type ListDepartmentsOptions struct {
	Page    uint   `url:"page,omitempty"`
	PerPage uint   `url:"per_page,omitempty"`
	Code    string `url:"code,omitempty"`
}

func (opt *ListDepartmentsOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListDepartments(ctx context.Context, opt *ListDepartmentsOptions) ([]Department, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/departments", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var departments []Department
	if err := decodeBody(res, &departments); err != nil {
		return nil, errors.WithStack(err)
	}
	return departments, nil
}

func (c *Client) CreateDepartment(ctx context.Context, department *Department) error {
	if department == nil {
		return nil
	}
	if err := validator.Struct(department); err != nil {
		return errors.WithStack(err)
	}
	b, err := json.Marshal(department)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.post(ctx, "/v1/departments", bytes.NewBuffer(b))
	return errors.WithStack(err)
}

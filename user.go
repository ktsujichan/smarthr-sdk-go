package smarthr

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type User struct {
	ID                   string               `json:"id" validate:"uuid"`
	Email                string               `json:"email" validate:"email"`
	Admin                bool                 `json:"admin"`
	Role                 *Role                `json:"role,omitempty"`
	CrewID               string               `json:"crew_id,omitempty" validate:"uuid"`
	Crew                 *Crew                `json:"crew,omitempty"`
	Tenants              []Tenant             `json:"tenants,omitempty"`
	InvitationCreatedAt  string               `json:"invitation_created_at,omitempty"`
	InvitationOpenedAt   string               `json:"invitation_opened_at,omitempty"`
	InvitationAcceptedAt string               `json:"invitation_accepted_at,omitempty"`
	InvitationAnsweredAt string               `json:"invitation_answered_at,omitempty"`
	SuppressedEmailLogs  []SuppressedEmailLog `json:"suppressed_email_logs,omitempty" validate:"omitempty,eq=bounced|eq=spam_reported|eq=blocked|eq=malformed"`
	HasPassword          bool                 `json:"has_password"`
	UpdatedAt            string               `json:"updated_at,omitempty"`
	CreatedAt            string               `json:"created_at,omitempty"`
}

func (c *Client) GetUser(ctx context.Context, id string) (*User, error) {
	p := fmt.Sprintf("/v1/users/%s", id)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var user User
	if err := decodeBody(res, &user); err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

type ListUsersOptions struct {
	Page    uint   `url:"page,omitempty"`
	PerPage uint   `url:"per_page,omitempty"`
	Embed   string `url:"embed,omitempty" validate:"omitempty,eq=crew"`
}

func (opt *ListUsersOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListUsers(ctx context.Context, opt *ListUsersOptions) ([]User, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/users", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var users []User
	if err := decodeBody(res, &users); err != nil {
		return nil, errors.WithStack(err)
	}
	return users, nil
}

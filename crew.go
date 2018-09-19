package smarthr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Crew struct {
	ID                     string            `json:"id" validate:"omitempty,uuid"`
	UserID                 string            `json:"user_id,omitempty"`
	BizEstablishmentID     string            `json:"biz_establishment_id"`
	EmpType                string            `json:"emp_type,omitempty" validate:"omitempty,emp_type"`
	EmploymentType         *EmploymentType   `json:"employment_type,omitempty"`
	EmpStatus              string            `json:"emp_status" validate:"omitempty,emp_status"`
	EmpCode                string            `json:"emp_code,omitempty"`
	LastName               string            `json:"last_name"`
	FirstName              string            `json:"first_name"`
	LastNameYomi           string            `json:"last_name_yomi"`
	FirstNameYomi          string            `json:"first_name_yomi"`
	BusinessLastName       string            `json:"business_last_name,omitempty"`
	BusinessFirstName      string            `json:"business_first_name,omitempty"`
	BusinessLastNameYomi   string            `json:"business_last_name_yomi,omitempty"`
	BusinessFirstNameYomi  string            `json:"business_first_name_yomi,omitempty"`
	BirthAt                string            `json:"birth_at,omitempty"`
	Gender                 string            `json:"gender" validate:"omitempty,gender"`
	TelNumber              string            `json:"tel_number,omitempty"`
	Email                  string            `json:"email,omitempty" validate:"omitempty,email"`
	Address                *Address          `json:"address,omitempty"`
	ResidentCardAddress    *Address          `json:"resident_card_address,omitempty"`
	ProfileImages          []Image           `json:"profile_images,omitempty"`
	EmergencyAddress       *Address          `json:"emergency_address,omitempty"`
	EmergencyRelationName  string            `json:"emergency_relation_name,omitempty"`
	EmergencyLastName      string            `json:"emergency_last_name,omitempty"`
	EmergencyFirstName     string            `json:"emergency_first_name,omitempty"`
	EmergencyLastNameYomi  string            `json:"emergency_last_name_yomi,omitempty"`
	EmergencyFirstNameYomi string            `json:"emergency_first_name_yomi,omitempty"`
	EmergencyTelNumber     string            `json:"emergency_tel_number,omitempty"`
	Department             string            `json:"department,omitempty"`
	Departments            []Department      `json:"departments,omitempty"`
	Position               string            `json:"position,omitempty"`
	Occupation             string            `json:"occupation,omitempty"`
	EnteredAt              string            `json:"entered_at,omitempty"`
	ResignedAt             string            `json:"resigned_at,omitempty"`
	ResignedReason         string            `json:"resigned_reason,omitempty"`
	Resume1                *Attachment       `json:"resume1,omitempty"`
	Resume2                *Attachment       `json:"resume2,omitempty"`
	BankAccounts           []BankAccount     `json:"bank_accounts,omitempty"`
	NearestStationAndLine  string            `json:"nearest_station_and_line,omitempty"`
	Commutation1Expenses   int               `json:"commutation_1_expenses,omitempty"`
	Commutation1Period     string            `json:"commutation_1_period,omitempty" validate:"omitempty,commutation_period"`
	Commutation1SingleFare int               `json:"commutation_1_single_fare,omitempty"`
	Commutation2Expenses   int               `json:"commutation_2_expenses,omitempty"`
	Commutation2Period     string            `json:"commutation_2_period,omitempty" validate:"omitempty,commutation_period"`
	Commutation2SingleFare int               `json:"commutation_2_single_fare,omitempty"`
	HavingSpouse           bool              `json:"having_spouse,omitempty"`
	CustomFields           []CrewCustomField `json:"custom_fields,omitempty"`
	UpdatedAt              string            `json:"updated_at"`
	CreatedAt              string            `json:"created_at"`
}

func (c *Crew) CustomFieldByTemplateID(templateID string) *CrewCustomField {
	for _, customFiled := range c.CustomFields {
		if customFiled.Template == nil {
			continue
		}
		if customFiled.Template.ID == templateID {
			return &customFiled
		}
	}
	return nil
}

type InviteCrewOptions struct {
	InviterUserID   string `json:"inviter_user_id" validate:"required"`
	CrewInputFormID string `json:"crew_input_form_id"`
}

func (opt *InviteCrewOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) InviteCrew(ctx context.Context, id string, opt *InviteCrewOptions) error {
	if err := opt.Validate(); err != nil {
		return errors.WithStack(err)
	}
	b, err := json.Marshal(opt)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.put(ctx, fmt.Sprintf("/v1/crews/%s/invite", id), bytes.NewBuffer(b))
	return errors.WithStack(err)
}

func (c *Client) DeleteCrew(ctx context.Context, id string) error {
	p := fmt.Sprintf("/v1/crews/%s", id)
	_, err := c.delete(ctx, p)
	return errors.WithStack(err)
}

func (c *Client) GetCrew(ctx context.Context, id string) (*Crew, error) {
	res, err := c.get(ctx, fmt.Sprintf("/v1/crews/%s", id), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var crew Crew
	if err := decodeBody(res, &crew); err != nil {
		return nil, errors.WithStack(err)
	}
	return &crew, nil
}

func (c *Client) UpdateCrew(ctx context.Context, id string, data map[string]interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.patch(ctx, fmt.Sprintf("/v1/crews/%s", id), bytes.NewBuffer(b))
	return errors.WithStack(err)
}

type ListCrewsOptions struct {
	Page          uint     `url:"page,omitempty"`
	PerPage       uint     `url:"per_page,omitempty"`
	EmpCode       string   `url:"emp_code,omitempty"`
	EmpType       string   `url:"emp_type,omitempty"`
	EmpStatus     string   `url:"emp_status,omitempty"`
	Gender        string   `url:"gender,omitempty" validate:"omitempty,eq=male|eq=female"`
	Sort          string   `url:"sort,omitempty"`
	EnteredBefore string   `url:"entered_before,omitempty"`
	EnteredAfter  string   `url:"entered_after,omitempty"`
	Query         string   `url:"q,omitempty"`
	Fields        []string `url:"fields,omitempty,comma"`
}

func (opt *ListCrewsOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListCrews(ctx context.Context, opt *ListCrewsOptions) ([]Crew, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/crews", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var crews []Crew
	if err := decodeBody(res, &crews); err != nil {
		return nil, errors.WithStack(err)
	}
	return crews, nil
}

func (c *Client) CreateCrew(ctx context.Context, crew *Crew) error {
	if crew == nil {
		return nil
	}
	if err := validator.Struct(crew); err != nil {
		return errors.WithStack(err)
	}
	b, err := json.Marshal(crew)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = c.post(ctx, "/v1/crews", bytes.NewBuffer(b))
	return errors.WithStack(err)
}

package smarthr

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type CrewInputForm struct {
	ID                               string                   `json:"id"`
	PresetType                       string                   `json:"preset_type,omitempty"`
	FormType                         string                   `json:"form_type"`
	Name                             string                   `json:"name"`
	WithMyNumber                     bool                     `json:"with_my_number,omitempty"`
	CrewMyNumberRequired             bool                     `json:"crew_my_number_required,omitempty"`
	DependentsMyNumberRequired       bool                     `json:"dependents_my_number_required,omitempty"`
	CrewMyNumberCardRequired         bool                     `json:"crew_my_number_card_required,omitempty"`
	CrewIdentificationRequired       bool                     `json:"crew_identification_required,omitempty"`
	CrewIsForHelIns                  bool                     `json:"crew_is_for_hel_ins,omitempty"`
	CrewIsForEmpIns                  bool                     `json:"crew_is_for_emp_ins,omitempty"`
	CrewIsForAccIns                  bool                     `json:"crew_is_for_acc_ins,omitempty"`
	CrewIsForTaxDeduction            bool                     `json:"crew_is_for_tax_deduction,omitempty"`
	CrewIsForShareholding            bool                     `json:"crew_is_for_shareholding,omitempty"`
	DependentsMyNumberCardRequired   bool                     `json:"dependents_my_number_card_required,omitempty"`
	DependentsIdentificationRequired bool                     `json:"dependents_identification_required,omitempty"`
	DependentsIsForHelIns            bool                     `json:"dependents_is_for_hel_ins,omitempty"`
	DependentsIsForEmpIns            bool                     `json:"dependents_is_for_emp_ins,omitempty"`
	DependentsIsForAccIns            bool                     `json:"dependents_is_for_acc_ins,omitempty"`
	DependentsIsForTaxDeduction      bool                     `json:"dependents_is_for_tax_deduction,omitempty"`
	DependentsIsForShareholding      bool                     `json:"dependents_is_for_shareholding,omitempty"`
	DependentsIsForCat3Ins           bool                     `json:"dependents_is_for_cat3_ins,omitempty"`
	FieldGroups                      *CrewInputFormFieldGroup `json:"field_groups,omitempty"`
	MailFormatID                     string                   `json:"mail_format_id,omitempty"`
	MailFormat                       *MailFormat              `json:"mail_format,omitempty"`
	UpdatedAt                        string                   `json:"updated_at"`
	CreatedAt                        string                   `json:"created_at"`
}

type CrewInputFormFieldGroup struct {
	BasicFieldGroupType            string                        `json:"basic_field_group_type,omitempty"`
	Hint                           string                        `json:"hint,omitempty"`
	Position                       int                           `json:"position,omitempty"`
	Fields                         *CrewInputFormField           `json:"fields,omitempty"`
	CrewCustomFieldTemplateGroupID string                        `json:"custom_field_template_group_id,omitempty"`
	CrewCustomFieldTemplateGroup   *CrewCustomFieldTemplateGroup `json:"custom_field_template_group,omitempty"`
}

type CrewInputFormField struct {
	AttributeName         string                   `json:"attribute_name,omitempty"`
	DisplayType           string                   `json:"display_type,omitempty"`
	CustomFieldTemplateID int                      `json:"custom_field_template_id,omitempty"`
	CustomFieldTemplate   *CrewCustomFieldTemplate `json:"custom_field_template,omitempty"`
}

func (c *Client) GetCrewInputForm(ctx context.Context, id string) (*CrewInputForm, error) {
	p := fmt.Sprintf("/v1/crew_input_forms/%s", id)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var form CrewInputForm
	if err := decodeBody(res, &form); err != nil {
		return nil, errors.WithStack(err)
	}
	return &form, nil
}

type ListCrewInputFormsOptions struct {
	Page    uint   `url:"page,omitempty"`
	PerPage uint   `url:"per_page,omitempty"`
	Embed   string `url:"embed,omitempty" validate:"omitempty,eq=mail_format"`
}

func (opt *ListCrewInputFormsOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListCrewInputForms(ctx context.Context, opt *ListCrewInputFormsOptions) ([]CrewInputForm, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/crew_input_forms", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var forms []CrewInputForm
	if err := decodeBody(res, &forms); err != nil {
		return nil, errors.WithStack(err)
	}
	return forms, nil
}

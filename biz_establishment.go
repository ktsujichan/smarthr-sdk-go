package smarthr

import (
	"context"

	"github.com/pkg/errors"
)

type BizEstablishment struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	SocInsName         string   `json:"soc_ins_name"`
	SocInsOwnerID      string   `json:"soc_ins_owner_id,omitempty"`
	SocInsOwner        *Crew    `json:"soc_ins_owner,omitempty"`
	SocInsAddress      *Address `json:"soc_ins_address,omitempty"`
	SocInsTelNumber    string   `json:"soc_ins_tel_number,omitempty"`
	LabInsName         string   `json:"lab_ins_name"`
	LabInsOwnerID      string   `json:"lab_ins_owner_id,omitempty"`
	LabInsOwner        *Crew    `json:"lab_ins_owner,omitempty"`
	LabInsAddress      *Address `json:"lab_ins_address,omitempty"`
	LabInsTelNumber    string   `json:"lab_ins_tel_number,omitempty"`
	JurisdictionTax    string   `json:"jurisdiction_tax,omitempty"`
	SalaryPayerAddress *Address `json:"salary_payer_address,omitempty"`
	UpdatedAt          string   `json:"updated_at,omitempty"`
	CreatedAt          string   `json:"created_at,omitempty"`
}

type ListBizEstablishmentsOptions struct {
	Page    uint   `url:"page,omitempty"`
	PerPage uint   `url:"per_page,omitempty"`
	Embed   string `url:"embed,omitempty" validate:"omitempty,eq=soc_ins_owner|eq=lab_ins_owner"`
}

func (opt *ListBizEstablishmentsOptions) Validate() error {
	if opt == nil {
		return nil
	}
	err := validator.Struct(opt)
	return errors.WithStack(err)
}

func (c *Client) ListBizEstablishments(ctx context.Context, opt *ListBizEstablishmentsOptions) ([]BizEstablishment, error) {
	if err := opt.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.get(ctx, "/v1/biz_establishments", opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var bizEstablishments []BizEstablishment
	if err := decodeBody(res, &bizEstablishments); err != nil {
		return nil, errors.WithStack(err)
	}
	return bizEstablishments, nil
}

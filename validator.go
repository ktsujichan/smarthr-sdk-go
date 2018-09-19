package smarthr

import (
	v "gopkg.in/go-playground/validator.v9"
)

var validator *v.Validate

func init() {
	validator = v.New()
	validator.RegisterAlias("commutation_period", "omitempty,eq=commutation_period_1_month|eq=commutation_period_3_month|eq=commutation_period_6_month")
	validator.RegisterAlias("crew_custom_field_template_type", "eq=date|eq=decimal|eq=enum|eq=file|eq=string|eq=text")
	validator.RegisterAlias("emp_status", "eq=employed|eq=absent|eq=retired")
	validator.RegisterAlias("emp_type", "eq=board_member|eq=full_timer|eq=contract_worker|eq=permatemp|eq=part_timer|eq=outsourcing_contractor|eq=etc")
	validator.RegisterAlias("gender", "eq=male|eq=female")
}

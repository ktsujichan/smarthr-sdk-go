package smarthr

type Tenant struct {
	ID                string     `json:"id" validate:"uuid"`
	Name              string     `json:"name,omitempty"`
	BaseTalent        bool       `json:"base_talent,omitempty"`
	Subdomains        *Subdomain `json:"subdomains,omitempty"`
	TrialStartAt      string     `json:"trial_start_at,omitempty"`
	TrialEndAt        string     `json:"trial_end_at,omitempty"`
	AddonSubscribable bool       `json:"addon_subscribable,omitempty"`
	UpdatedAt         string     `json:"updated_at,omitempty"`
	CreatedAt         string     `json:"created_at,omitempty"`
}

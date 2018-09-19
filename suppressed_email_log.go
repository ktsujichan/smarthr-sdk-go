package smarthr

type SuppressedEmailLog struct {
	ID              string `json:"id" validate:"uuid"`
	SuppressionType int    `json:"suppression_type,omitempty"`
	Reason          string `json:"reason,omitempty"`
	SuppressedAt    string `json:"suppressed_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
}

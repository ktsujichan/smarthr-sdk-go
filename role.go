package smarthr

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CrewsScope  string `json:"crews_scope"`
}

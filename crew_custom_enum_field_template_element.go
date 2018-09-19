package smarthr

type CrewCustomEnumFieldTemplateElement struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhysicalName string `json:"physical_name,omitempty"`
	Position     int    `json:"position,omitempty"`
}

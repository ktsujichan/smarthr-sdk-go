package smarthr

type CrewCustomField struct {
	Value    interface{}              `json:"value,omitempty"`
	Template *CrewCustomFieldTemplate `json:"template,omitempty"`
}

func (c *CrewCustomField) String() string {
	if c.Template == nil || c.Value == nil {
		return ""
	}
	switch c.Template.Type {
	case "string", "text", "date", "file":
		if v, ok := c.Value.(string); ok {
			return v
		}
	}
	return ""
}

func (c *CrewCustomField) Float64() float64 {
	if c.Template == nil || c.Value == nil {
		return 0
	}
	switch c.Template.Type {
	case "decimal":
		if v, ok := c.Value.(float64); ok {
			return v
		}
	}
	return 0
}

func (c *CrewCustomField) Int64() int64 {
	if c.Template == nil || c.Value == nil {
		return 0
	}
	switch c.Template.Type {
	case "decimal":
		if v, ok := c.Value.(int64); ok {
			return v
		}
	}
	return 0
}

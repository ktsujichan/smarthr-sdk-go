package smarthr

type Image struct {
	SizeType string `json:"size_type,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
	URL      string `json:"url,omitempty"`
}

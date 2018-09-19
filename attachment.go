package smarthr

type Attachment struct {
	FileName string `json:"file_name,omitempty"`
	URL      string `json:"url,omitempty"`
}

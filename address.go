package smarthr

type Address struct {
	CountryNumber int    `json:"country_number,omitempty"`
	ZipCode       string `json:"zip_code,omitempty"`
	Pref          string `json:"pref,omitempty"`
	City          string `json:"city"`
	Street        string `json:"street,omitempty"`
	Building      string `json:"building,omitempty"`
	LiteralYomi   string `json:"literal_yomi,omitempty"`
}

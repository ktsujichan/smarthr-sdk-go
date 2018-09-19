package smarthr

import (
	"encoding/json"
)

type Error struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func (e Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

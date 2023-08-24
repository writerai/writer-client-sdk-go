// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type FailMessage struct {
	Description string      `json:"description"`
	Extras      interface{} `json:"extras"`
	Key         string      `json:"key"`
}

var _ error = &FailMessage{}

func (e *FailMessage) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

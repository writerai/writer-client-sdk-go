// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

// FailResponse - Bad Request
type FailResponse struct {
	Errors []shared.FailMessage `json:"errors,omitempty"`
	Extras interface{}          `json:"extras"`
	Tpe    string               `json:"tpe"`
}

var _ error = &FailResponse{}

func (e *FailResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreateTermsRequestFailHandling string

const (
	CreateTermsRequestFailHandlingAccumulate   CreateTermsRequestFailHandling = "accumulate"
	CreateTermsRequestFailHandlingValidate     CreateTermsRequestFailHandling = "validate"
	CreateTermsRequestFailHandlingSkip         CreateTermsRequestFailHandling = "skip"
	CreateTermsRequestFailHandlingValidateOnly CreateTermsRequestFailHandling = "validateOnly"
)

func (e CreateTermsRequestFailHandling) ToPointer() *CreateTermsRequestFailHandling {
	return &e
}

func (e *CreateTermsRequestFailHandling) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "accumulate":
		fallthrough
	case "validate":
		fallthrough
	case "skip":
		fallthrough
	case "validateOnly":
		*e = CreateTermsRequestFailHandling(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTermsRequestFailHandling: %v", v)
	}
}

type CreateTermsRequest struct {
	FailHandling *CreateTermsRequestFailHandling `json:"failHandling,omitempty"`
	Models       []TermCreate                    `json:"models,omitempty"`
}

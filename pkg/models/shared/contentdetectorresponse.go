// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Label string

const (
	LabelFake Label = "fake"
	LabelReal Label = "real"
)

func (e Label) ToPointer() *Label {
	return &e
}

func (e *Label) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		fallthrough
	case "real":
		*e = Label(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Label: %v", v)
	}
}

type ContentDetectorResponse struct {
	Label Label   `json:"label"`
	Score float64 `json:"score"`
}

func (o *ContentDetectorResponse) GetLabel() Label {
	if o == nil {
		return Label("")
	}
	return o.Label
}

func (o *ContentDetectorResponse) GetScore() float64 {
	if o == nil {
		return 0.0
	}
	return o.Score
}

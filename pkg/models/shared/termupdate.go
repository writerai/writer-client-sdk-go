// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TermUpdatePos string

const (
	TermUpdatePosNoun      TermUpdatePos = "noun"
	TermUpdatePosVerb      TermUpdatePos = "verb"
	TermUpdatePosAdverb    TermUpdatePos = "adverb"
	TermUpdatePosAdjective TermUpdatePos = "adjective"
)

func (e TermUpdatePos) ToPointer() *TermUpdatePos {
	return &e
}

func (e *TermUpdatePos) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "noun":
		fallthrough
	case "verb":
		fallthrough
	case "adverb":
		fallthrough
	case "adjective":
		*e = TermUpdatePos(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TermUpdatePos: %v", v)
	}
}

type TermUpdateType string

const (
	TermUpdateTypeApproved TermUpdateType = "approved"
	TermUpdateTypeBanned   TermUpdateType = "banned"
	TermUpdateTypePending  TermUpdateType = "pending"
)

func (e TermUpdateType) ToPointer() *TermUpdateType {
	return &e
}

func (e *TermUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "approved":
		fallthrough
	case "banned":
		fallthrough
	case "pending":
		*e = TermUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TermUpdateType: %v", v)
	}
}

type TermUpdate struct {
	ApprovedTermExtension *ApprovedTermExtensionCreate `json:"approvedTermExtension,omitempty"`
	CaseSensitive         bool                         `json:"caseSensitive"`
	Description           *string                      `json:"description,omitempty"`
	Examples              []TermExampleCreate          `json:"examples,omitempty"`
	Highlight             *bool                        `json:"highlight,omitempty"`
	ID                    int64                        `json:"id"`
	LinkedTerms           []LinkedTermCreate           `json:"linkedTerms,omitempty"`
	Mistakes              []TermMistakeCreate          `json:"mistakes,omitempty"`
	Pos                   *TermUpdatePos               `json:"pos,omitempty"`
	Tags                  []TermTagCreate              `json:"tags,omitempty"`
	Term                  string                       `json:"term"`
	Type                  TermUpdateType               `json:"type"`
}

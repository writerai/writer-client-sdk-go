package shared

type TermCreatePosEnum string

const (
	TermCreatePosEnumNoun      TermCreatePosEnum = "noun"
	TermCreatePosEnumVerb      TermCreatePosEnum = "verb"
	TermCreatePosEnumAdverb    TermCreatePosEnum = "adverb"
	TermCreatePosEnumAdjective TermCreatePosEnum = "adjective"
)

type TermCreateTypeEnum string

const (
	TermCreateTypeEnumApproved TermCreateTypeEnum = "approved"
	TermCreateTypeEnumBanned   TermCreateTypeEnum = "banned"
	TermCreateTypeEnumPending  TermCreateTypeEnum = "pending"
)

type TermCreate struct {
	ApprovedTermExtension *ApprovedTermExtensionCreate `json:"approvedTermExtension,omitempty"`
	CaseSensitive         bool                         `json:"caseSensitive"`
	Description           *string                      `json:"description,omitempty"`
	Examples              []TermExampleCreate          `json:"examples,omitempty"`
	Highlight             *bool                        `json:"highlight,omitempty"`
	LinkedTerms           []LinkedTermCreate           `json:"linkedTerms,omitempty"`
	Mistakes              []TermMistakeCreate          `json:"mistakes,omitempty"`
	Pos                   *TermCreatePosEnum           `json:"pos,omitempty"`
	Reference             *string                      `json:"reference,omitempty"`
	Tags                  []TermTagCreate              `json:"tags,omitempty"`
	Term                  string                       `json:"term"`
	Type                  TermCreateTypeEnum           `json:"type"`
}

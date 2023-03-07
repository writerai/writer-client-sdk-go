package shared

type TermUpdatePosEnum string

const (
	TermUpdatePosEnumNoun      TermUpdatePosEnum = "noun"
	TermUpdatePosEnumVerb      TermUpdatePosEnum = "verb"
	TermUpdatePosEnumAdverb    TermUpdatePosEnum = "adverb"
	TermUpdatePosEnumAdjective TermUpdatePosEnum = "adjective"
)

type TermUpdateTypeEnum string

const (
	TermUpdateTypeEnumApproved TermUpdateTypeEnum = "approved"
	TermUpdateTypeEnumBanned   TermUpdateTypeEnum = "banned"
	TermUpdateTypeEnumPending  TermUpdateTypeEnum = "pending"
)

type TermUpdate struct {
	ApprovedTermExtension *ApprovedTermExtensionCreate `json:"approvedTermExtension,omitempty"`
	CaseSensitive         bool                         `json:"caseSensitive"`
	Description           *string                      `json:"description,omitempty"`
	Examples              []TermExampleCreate          `json:"examples,omitempty"`
	Highlight             *bool                        `json:"highlight,omitempty"`
	ID                    int64                        `json:"id"`
	LinkedTerms           []LinkedTermCreate           `json:"linkedTerms,omitempty"`
	Mistakes              []TermMistakeCreate          `json:"mistakes,omitempty"`
	Pos                   *TermUpdatePosEnum           `json:"pos,omitempty"`
	Tags                  []TermTagCreate              `json:"tags,omitempty"`
	Term                  string                       `json:"term"`
	Type                  TermUpdateTypeEnum           `json:"type"`
}

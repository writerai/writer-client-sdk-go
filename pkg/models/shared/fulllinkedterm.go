package shared

type FullLinkedTermPosEnum string

const (
	FullLinkedTermPosEnumNoun      FullLinkedTermPosEnum = "noun"
	FullLinkedTermPosEnumVerb      FullLinkedTermPosEnum = "verb"
	FullLinkedTermPosEnumAdverb    FullLinkedTermPosEnum = "adverb"
	FullLinkedTermPosEnumAdjective FullLinkedTermPosEnum = "adjective"
)

type FullLinkedTerm struct {
	ApprovedTermExtension *ApprovedTermExtension `json:"approvedTermExtension,omitempty"`
	CaseSensitive         bool                   `json:"caseSensitive"`
	ID                    *int64                 `json:"id,omitempty"`
	LinkedTermID          int64                  `json:"linkedTermId"`
	Pos                   *FullLinkedTermPosEnum `json:"pos,omitempty"`
	Term                  string                 `json:"term"`
	TermID                int64                  `json:"termId"`
}

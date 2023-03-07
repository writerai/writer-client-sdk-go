package shared

type TermMistakeCreatePosEnum string

const (
	TermMistakeCreatePosEnumNoun      TermMistakeCreatePosEnum = "noun"
	TermMistakeCreatePosEnumVerb      TermMistakeCreatePosEnum = "verb"
	TermMistakeCreatePosEnumAdverb    TermMistakeCreatePosEnum = "adverb"
	TermMistakeCreatePosEnumAdjective TermMistakeCreatePosEnum = "adjective"
)

type TermMistakeCreate struct {
	CaseSensitive bool                      `json:"caseSensitive"`
	Mistake       string                    `json:"mistake"`
	Pos           *TermMistakeCreatePosEnum `json:"pos,omitempty"`
	Reference     *string                   `json:"reference,omitempty"`
}

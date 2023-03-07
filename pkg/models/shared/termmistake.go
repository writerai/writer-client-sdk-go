package shared

type TermMistakePosEnum string

const (
	TermMistakePosEnumNoun      TermMistakePosEnum = "noun"
	TermMistakePosEnumVerb      TermMistakePosEnum = "verb"
	TermMistakePosEnumAdverb    TermMistakePosEnum = "adverb"
	TermMistakePosEnumAdjective TermMistakePosEnum = "adjective"
)

type TermMistake struct {
	CaseSensitive bool                `json:"caseSensitive"`
	ID            *int64              `json:"id,omitempty"`
	Mistake       string              `json:"mistake"`
	Pos           *TermMistakePosEnum `json:"pos,omitempty"`
	TermBankID    int64               `json:"termBankId"`
	TermID        int64               `json:"termId"`
}

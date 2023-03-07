package shared

type TermExampleTypeEnum string

const (
	TermExampleTypeEnumGood TermExampleTypeEnum = "good"
	TermExampleTypeEnumBad  TermExampleTypeEnum = "bad"
)

type TermExample struct {
	Example    string              `json:"example"`
	ID         *int64              `json:"id,omitempty"`
	TermBankID int64               `json:"termBankId"`
	TermID     int64               `json:"termId"`
	Type       TermExampleTypeEnum `json:"type"`
}

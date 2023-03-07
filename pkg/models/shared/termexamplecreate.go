package shared

type TermExampleCreateTypeEnum string

const (
	TermExampleCreateTypeEnumGood TermExampleCreateTypeEnum = "good"
	TermExampleCreateTypeEnumBad  TermExampleCreateTypeEnum = "bad"
)

type TermExampleCreate struct {
	Example string                    `json:"example"`
	Type    TermExampleCreateTypeEnum `json:"type"`
}

package shared

type ApprovedTermExtensionCreate struct {
	Capitalize        bool `json:"capitalize"`
	FixCase           bool `json:"fixCase"`
	FixCommonMistakes bool `json:"fixCommonMistakes"`
}

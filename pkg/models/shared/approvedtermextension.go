package shared

type ApprovedTermExtension struct {
	Capitalize        bool   `json:"capitalize"`
	FixCase           bool   `json:"fixCase"`
	FixCommonMistakes bool   `json:"fixCommonMistakes"`
	ID                *int64 `json:"id,omitempty"`
	TermID            int64  `json:"termId"`
}

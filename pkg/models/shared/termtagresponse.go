package shared

type TermTagResponse struct {
	ID          int64  `json:"id"`
	ParentTagID int64  `json:"parentTagId"`
	Tag         string `json:"tag"`
	TermID      int64  `json:"termId"`
}

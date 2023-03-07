package shared

type ContentRequest struct {
	Content  string          `json:"content"`
	Settings ContentSettings `json:"settings"`
}

package shared

type LinkedTermCreate struct {
	LinkedTermID *int64  `json:"linkedTermId,omitempty"`
	Reference    *string `json:"reference,omitempty"`
}

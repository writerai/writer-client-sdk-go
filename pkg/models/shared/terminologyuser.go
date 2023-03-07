package shared

type TerminologyUser struct {
	Email    *string `json:"email,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	ID       int64   `json:"id"`
}

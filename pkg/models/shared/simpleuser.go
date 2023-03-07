package shared

type SimpleUser struct {
	Email     *string `json:"email,omitempty"`
	FirstName string  `json:"firstName"`
	ID        int64   `json:"id"`
	LastName  *string `json:"lastName,omitempty"`
}

package shared

type MagicRequestInput struct {
	Name  string   `json:"name"`
	Value []string `json:"value,omitempty"`
}

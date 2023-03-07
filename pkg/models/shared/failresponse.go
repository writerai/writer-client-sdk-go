package shared

type FailResponse struct {
	Errors []FailMessage `json:"errors,omitempty"`
	Extras interface{}   `json:"extras"`
	Tpe    string        `json:"tpe"`
}

package shared

type FailMessage struct {
	Description string      `json:"description"`
	Extras      interface{} `json:"extras"`
	Key         string      `json:"key"`
}

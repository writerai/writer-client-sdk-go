package shared

type CreateTermsResponse struct {
	Fails  []FailMessage      `json:"fails,omitempty"`
	Models []FullTermWithUser `json:"models,omitempty"`
}

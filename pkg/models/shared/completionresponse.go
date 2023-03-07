package shared

type CompletionResponse struct {
	Choices []CompletionGenerationChoice `json:"choices,omitempty"`
}

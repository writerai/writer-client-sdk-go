package shared

type CompletionGenerationChoice struct {
	Logprobs *CompletionGenerationChoiceLogprobs `json:"logprobs,omitempty"`
	Text     string                              `json:"text"`
}

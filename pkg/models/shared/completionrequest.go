package shared

type CompletionRequest struct {
	BestOf           *int64   `json:"bestOf,omitempty"`
	FrequencyPenalty *float64 `json:"frequencyPenalty,omitempty"`
	Logprobs         *int64   `json:"logprobs,omitempty"`
	MaxTokens        *int64   `json:"maxTokens,omitempty"`
	MinTokens        *int64   `json:"minTokens,omitempty"`
	N                *int64   `json:"n,omitempty"`
	PresencePenalty  *float64 `json:"presencePenalty,omitempty"`
	Prompt           string   `json:"prompt"`
	Stop             []string `json:"stop,omitempty"`
	Temperature      *float64 `json:"temperature,omitempty"`
	TopP             *float64 `json:"topP,omitempty"`
}

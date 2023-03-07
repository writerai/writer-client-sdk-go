package shared

type CompletionGenerationChoiceLogprobs struct {
	TextOffset    []int64             `json:"textOffset,omitempty"`
	TokenLogprobs []float64           `json:"tokenLogprobs,omitempty"`
	Tokens        []string            `json:"tokens,omitempty"`
	TopLogprobs   []map[string]string `json:"topLogprobs,omitempty"`
}

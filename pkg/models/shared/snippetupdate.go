package shared

type SnippetUpdate struct {
	Description *string        `json:"description,omitempty"`
	ID          string         `json:"id"`
	Shortcut    *string        `json:"shortcut,omitempty"`
	Snippet     string         `json:"snippet"`
	Tags        []SnippetTagV2 `json:"tags,omitempty"`
}

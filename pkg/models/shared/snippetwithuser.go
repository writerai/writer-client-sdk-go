package shared

import (
	"time"
)

type SnippetWithUser struct {
	CreatedUser      TerminologyUser `json:"createdUser"`
	CreationTime     time.Time       `json:"creationTime"`
	Description      *string         `json:"description,omitempty"`
	ID               string          `json:"id"`
	ModificationTime time.Time       `json:"modificationTime"`
	ModifiedUser     TerminologyUser `json:"modifiedUser"`
	Shortcut         *string         `json:"shortcut,omitempty"`
	Snippet          string          `json:"snippet"`
	Tags             []SnippetTagV2  `json:"tags,omitempty"`
}

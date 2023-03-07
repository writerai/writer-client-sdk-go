package shared

import (
	"time"
)

type TemplateDetailsResponse struct {
	CategoryID       int64     `json:"categoryId"`
	CreationTime     time.Time `json:"creationTime"`
	Description      *string   `json:"description,omitempty"`
	GuideURL         *string   `json:"guideUrl,omitempty"`
	ID               string    `json:"id"`
	Inputs           []Input   `json:"inputs,omitempty"`
	ModificationTime time.Time `json:"modificationTime"`
	Name             string    `json:"name"`
}

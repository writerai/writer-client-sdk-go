package shared

import (
	"time"
)

type Draft struct {
	Body           string      `json:"body"`
	CreatedUserID  int64       `json:"createdUserId"`
	CreationTime   time.Time   `json:"creationTime"`
	Deleted        bool        `json:"deleted"`
	DocumentID     string      `json:"documentId"`
	ID             *int64      `json:"id,omitempty"`
	Inputs         interface{} `json:"inputs"`
	OrganizationID int64       `json:"organizationId"`
	TeamID         int64       `json:"teamId"`
	TemplateID     string      `json:"templateId"`
	Title          *string     `json:"title,omitempty"`
}

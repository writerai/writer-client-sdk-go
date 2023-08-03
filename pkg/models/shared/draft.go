// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

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

func (o *Draft) GetBody() string {
	if o == nil {
		return ""
	}
	return o.Body
}

func (o *Draft) GetCreatedUserID() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedUserID
}

func (o *Draft) GetCreationTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreationTime
}

func (o *Draft) GetDeleted() bool {
	if o == nil {
		return false
	}
	return o.Deleted
}

func (o *Draft) GetDocumentID() string {
	if o == nil {
		return ""
	}
	return o.DocumentID
}

func (o *Draft) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Draft) GetInputs() interface{} {
	if o == nil {
		return nil
	}
	return o.Inputs
}

func (o *Draft) GetOrganizationID() int64 {
	if o == nil {
		return 0
	}
	return o.OrganizationID
}

func (o *Draft) GetTeamID() int64 {
	if o == nil {
		return 0
	}
	return o.TeamID
}

func (o *Draft) GetTemplateID() string {
	if o == nil {
		return ""
	}
	return o.TemplateID
}

func (o *Draft) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

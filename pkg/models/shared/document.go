// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type DocumentAccess string

const (
	DocumentAccessPrivate DocumentAccess = "private"
	DocumentAccessPublic  DocumentAccess = "public"
	DocumentAccessShared  DocumentAccess = "shared"
)

func (e DocumentAccess) ToPointer() *DocumentAccess {
	return &e
}

func (e *DocumentAccess) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "private":
		fallthrough
	case "public":
		fallthrough
	case "shared":
		*e = DocumentAccess(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DocumentAccess: %v", v)
	}
}

type Document struct {
	Access           DocumentAccess `json:"access"`
	Content          string         `json:"content"`
	CreatedUser      *SimpleUser    `json:"createdUser,omitempty"`
	CreationTime     time.Time      `json:"creationTime"`
	ID               int64          `json:"id"`
	ModificationTime time.Time      `json:"modificationTime"`
	ModifiedUser     *SimpleUser    `json:"modifiedUser,omitempty"`
	OrganizationID   int64          `json:"organizationId"`
	Score            int64          `json:"score"`
	TeamID           int64          `json:"teamId"`
	Title            string         `json:"title"`
}

func (o *Document) GetAccess() DocumentAccess {
	if o == nil {
		return DocumentAccess("")
	}
	return o.Access
}

func (o *Document) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *Document) GetCreatedUser() *SimpleUser {
	if o == nil {
		return nil
	}
	return o.CreatedUser
}

func (o *Document) GetCreationTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreationTime
}

func (o *Document) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Document) GetModificationTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ModificationTime
}

func (o *Document) GetModifiedUser() *SimpleUser {
	if o == nil {
		return nil
	}
	return o.ModifiedUser
}

func (o *Document) GetOrganizationID() int64 {
	if o == nil {
		return 0
	}
	return o.OrganizationID
}

func (o *Document) GetScore() int64 {
	if o == nil {
		return 0
	}
	return o.Score
}

func (o *Document) GetTeamID() int64 {
	if o == nil {
		return 0
	}
	return o.TeamID
}

func (o *Document) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

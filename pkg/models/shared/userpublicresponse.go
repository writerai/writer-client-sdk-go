package shared

import (
	"time"
)

type UserPublicResponseAccountStatusEnum string

const (
	UserPublicResponseAccountStatusEnumInvited  UserPublicResponseAccountStatusEnum = "invited"
	UserPublicResponseAccountStatusEnumSignedUp UserPublicResponseAccountStatusEnum = "signed_up"
)

type UserPublicResponse struct {
	AccountStatus  UserPublicResponseAccountStatusEnum `json:"accountStatus"`
	Avatar         *string                             `json:"avatar,omitempty"`
	CreatedAt      time.Time                           `json:"createdAt"`
	Email          *string                             `json:"email,omitempty"`
	FirstName      string                              `json:"firstName"`
	FullName       string                              `json:"fullName"`
	ID             int64                               `json:"id"`
	InvitedBy      *int64                              `json:"invitedBy,omitempty"`
	LastName       *string                             `json:"lastName,omitempty"`
	LastSeenOnline *time.Time                          `json:"lastSeenOnline,omitempty"`
	Timezone       *string                             `json:"timezone,omitempty"`
}

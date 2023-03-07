package shared

import (
	"time"
)

type PageWithSectionResponseStatusEnum string

const (
	PageWithSectionResponseStatusEnumLive    PageWithSectionResponseStatusEnum = "live"
	PageWithSectionResponseStatusEnumOffline PageWithSectionResponseStatusEnum = "offline"
)

type PageWithSectionResponse struct {
	Content   *string                           `json:"content,omitempty"`
	CreatedAt time.Time                         `json:"createdAt"`
	ID        int64                             `json:"id"`
	Order     int64                             `json:"order"`
	Section   *SectionInfo                      `json:"section,omitempty"`
	Status    PageWithSectionResponseStatusEnum `json:"status"`
	Title     string                            `json:"title"`
	UpdatedAt time.Time                         `json:"updatedAt"`
	UpdatedBy *SimpleUser                       `json:"updatedBy,omitempty"`
	URL       string                            `json:"url"`
}

package shared

import (
	"time"
)

type PagePublicAPIResponseStatusEnum string

const (
	PagePublicAPIResponseStatusEnumLive    PagePublicAPIResponseStatusEnum = "live"
	PagePublicAPIResponseStatusEnumOffline PagePublicAPIResponseStatusEnum = "offline"
)

type PagePublicAPIResponse struct {
	CreatedAt time.Time                       `json:"createdAt"`
	ID        int64                           `json:"id"`
	Order     int64                           `json:"order"`
	Section   *SectionInfo                    `json:"section,omitempty"`
	Status    PagePublicAPIResponseStatusEnum `json:"status"`
	Title     string                          `json:"title"`
	UpdatedAt time.Time                       `json:"updatedAt"`
	UpdatedBy *SimpleUser                     `json:"updatedBy,omitempty"`
	URL       string                          `json:"url"`
}

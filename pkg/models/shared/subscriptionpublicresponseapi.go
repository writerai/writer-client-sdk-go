package shared

import (
	"time"
)

type SubscriptionPublicResponseAPIProductNameEnum string

const (
	SubscriptionPublicResponseAPIProductNameEnumFree       SubscriptionPublicResponseAPIProductNameEnum = "free"
	SubscriptionPublicResponseAPIProductNameEnumPro        SubscriptionPublicResponseAPIProductNameEnum = "pro"
	SubscriptionPublicResponseAPIProductNameEnumTeam       SubscriptionPublicResponseAPIProductNameEnum = "team"
	SubscriptionPublicResponseAPIProductNameEnumEnterprise SubscriptionPublicResponseAPIProductNameEnum = "enterprise"
	SubscriptionPublicResponseAPIProductNameEnumLegacy     SubscriptionPublicResponseAPIProductNameEnum = "legacy"
)

type SubscriptionPublicResponseAPIStatusEnum string

const (
	SubscriptionPublicResponseAPIStatusEnumTrialing          SubscriptionPublicResponseAPIStatusEnum = "trialing"
	SubscriptionPublicResponseAPIStatusEnumActive            SubscriptionPublicResponseAPIStatusEnum = "active"
	SubscriptionPublicResponseAPIStatusEnumPastDue           SubscriptionPublicResponseAPIStatusEnum = "past_due"
	SubscriptionPublicResponseAPIStatusEnumIncomplete        SubscriptionPublicResponseAPIStatusEnum = "incomplete"
	SubscriptionPublicResponseAPIStatusEnumIncompleteExpired SubscriptionPublicResponseAPIStatusEnum = "incomplete_expired"
	SubscriptionPublicResponseAPIStatusEnumUnpaid            SubscriptionPublicResponseAPIStatusEnum = "unpaid"
	SubscriptionPublicResponseAPIStatusEnumCanceled          SubscriptionPublicResponseAPIStatusEnum = "canceled"
)

type SubscriptionPublicResponseAPI struct {
	CreatedAt      time.Time                                    `json:"createdAt"`
	Meta           MetaData                                     `json:"meta"`
	ProductName    SubscriptionPublicResponseAPIProductNameEnum `json:"productName"`
	Seats          int64                                        `json:"seats"`
	Status         SubscriptionPublicResponseAPIStatusEnum      `json:"status"`
	SubscriptionID string                                       `json:"subscriptionId"`
	Usage          Usage                                        `json:"usage"`
}

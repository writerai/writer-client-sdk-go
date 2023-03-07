package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type FindTermsPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type FindTermsPartOfSpeechEnum string

const (
	FindTermsPartOfSpeechEnumNoun      FindTermsPartOfSpeechEnum = "noun"
	FindTermsPartOfSpeechEnumVerb      FindTermsPartOfSpeechEnum = "verb"
	FindTermsPartOfSpeechEnumAdverb    FindTermsPartOfSpeechEnum = "adverb"
	FindTermsPartOfSpeechEnumAdjective FindTermsPartOfSpeechEnum = "adjective"
)

type FindTermsSortFieldEnum string

const (
	FindTermsSortFieldEnumTerm             FindTermsSortFieldEnum = "term"
	FindTermsSortFieldEnumCreationTime     FindTermsSortFieldEnum = "creationTime"
	FindTermsSortFieldEnumModificationTime FindTermsSortFieldEnum = "modificationTime"
	FindTermsSortFieldEnumType             FindTermsSortFieldEnum = "type"
)

type FindTermsSortOrderEnum string

const (
	FindTermsSortOrderEnumAsc  FindTermsSortOrderEnum = "asc"
	FindTermsSortOrderEnumDesc FindTermsSortOrderEnum = "desc"
)

type FindTermsTypeEnum string

const (
	FindTermsTypeEnumApproved FindTermsTypeEnum = "approved"
	FindTermsTypeEnumBanned   FindTermsTypeEnum = "banned"
	FindTermsTypeEnumPending  FindTermsTypeEnum = "pending"
)

type FindTermsQueryParams struct {
	Limit        *int64                     `queryParam:"style=form,explode=true,name=limit"`
	Offset       *int64                     `queryParam:"style=form,explode=true,name=offset"`
	PartOfSpeech *FindTermsPartOfSpeechEnum `queryParam:"style=form,explode=true,name=partOfSpeech"`
	SortField    *FindTermsSortFieldEnum    `queryParam:"style=form,explode=true,name=sortField"`
	SortOrder    *FindTermsSortOrderEnum    `queryParam:"style=form,explode=true,name=sortOrder"`
	Tags         []string                   `queryParam:"style=form,explode=true,name=tags"`
	Term         *string                    `queryParam:"style=form,explode=true,name=term"`
	Type         *FindTermsTypeEnum         `queryParam:"style=form,explode=true,name=type"`
}

type FindTermsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type FindTermsRequest struct {
	PathParams  FindTermsPathParams
	QueryParams FindTermsQueryParams
	Headers     FindTermsHeaders
}

type FindTermsResponse struct {
	ContentType                     string
	FailResponse                    *shared.FailResponse
	Headers                         map[string][]string
	PaginatedResultFullTermWithUser *shared.PaginatedResultFullTermWithUser
	StatusCode                      int
	RawResponse                     *http.Response
}

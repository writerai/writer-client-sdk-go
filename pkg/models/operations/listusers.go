package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ListUsersSortFieldEnum string

const (
	ListUsersSortFieldEnumID               ListUsersSortFieldEnum = "id"
	ListUsersSortFieldEnumName             ListUsersSortFieldEnum = "name"
	ListUsersSortFieldEnumCreationTime     ListUsersSortFieldEnum = "creationTime"
	ListUsersSortFieldEnumDeleted          ListUsersSortFieldEnum = "deleted"
	ListUsersSortFieldEnumModificationTime ListUsersSortFieldEnum = "modificationTime"
	ListUsersSortFieldEnumEmail            ListUsersSortFieldEnum = "email"
	ListUsersSortFieldEnumLastSeen         ListUsersSortFieldEnum = "lastSeen"
)

type ListUsersSortOrderEnum string

const (
	ListUsersSortOrderEnumAsc  ListUsersSortOrderEnum = "asc"
	ListUsersSortOrderEnumDesc ListUsersSortOrderEnum = "desc"
)

type ListUsersQueryParams struct {
	Limit     *int64                  `queryParam:"style=form,explode=true,name=limit"`
	Offset    *int64                  `queryParam:"style=form,explode=true,name=offset"`
	Search    *string                 `queryParam:"style=form,explode=true,name=search"`
	SortField *ListUsersSortFieldEnum `queryParam:"style=form,explode=true,name=sortField"`
	SortOrder *ListUsersSortOrderEnum `queryParam:"style=form,explode=true,name=sortOrder"`
}

type ListUsersHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type ListUsersRequest struct {
	QueryParams ListUsersQueryParams
	Headers     ListUsersHeaders
}

type ListUsersResponse struct {
	ContentType                       string
	FailResponse                      *shared.FailResponse
	Headers                           map[string][]string
	PaginatedResultUserPublicResponse *shared.PaginatedResultUserPublicResponse
	StatusCode                        int
	RawResponse                       *http.Response
}

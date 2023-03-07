package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type FindSnippetsPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type FindSnippetsSortFieldEnum string

const (
	FindSnippetsSortFieldEnumShortcut         FindSnippetsSortFieldEnum = "shortcut"
	FindSnippetsSortFieldEnumCreationTime     FindSnippetsSortFieldEnum = "creationTime"
	FindSnippetsSortFieldEnumModificationTime FindSnippetsSortFieldEnum = "modificationTime"
)

type FindSnippetsSortOrderEnum string

const (
	FindSnippetsSortOrderEnumAsc  FindSnippetsSortOrderEnum = "asc"
	FindSnippetsSortOrderEnumDesc FindSnippetsSortOrderEnum = "desc"
)

type FindSnippetsQueryParams struct {
	Limit     *int64                     `queryParam:"style=form,explode=true,name=limit"`
	Offset    *int64                     `queryParam:"style=form,explode=true,name=offset"`
	Search    *string                    `queryParam:"style=form,explode=true,name=search"`
	Shortcuts []string                   `queryParam:"style=form,explode=true,name=shortcuts"`
	SortField *FindSnippetsSortFieldEnum `queryParam:"style=form,explode=true,name=sortField"`
	SortOrder *FindSnippetsSortOrderEnum `queryParam:"style=form,explode=true,name=sortOrder"`
	Tags      []string                   `queryParam:"style=form,explode=true,name=tags"`
}

type FindSnippetsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type FindSnippetsRequest struct {
	PathParams  FindSnippetsPathParams
	QueryParams FindSnippetsQueryParams
	Headers     FindSnippetsHeaders
}

type FindSnippetsResponse struct {
	ContentType                    string
	FailResponse                   *shared.FailResponse
	Headers                        map[string][]string
	PaginatedResultSnippetWithUser *shared.PaginatedResultSnippetWithUser
	StatusCode                     int
	RawResponse                    *http.Response
}

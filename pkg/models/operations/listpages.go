package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ListPagesStatusEnum string

const (
	ListPagesStatusEnumLive    ListPagesStatusEnum = "live"
	ListPagesStatusEnumOffline ListPagesStatusEnum = "offline"
)

type ListPagesQueryParams struct {
	Limit  *int64               `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64               `queryParam:"style=form,explode=true,name=offset"`
	Status *ListPagesStatusEnum `queryParam:"style=form,explode=true,name=status"`
}

type ListPagesHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type ListPagesRequest struct {
	QueryParams ListPagesQueryParams
	Headers     ListPagesHeaders
}

type ListPagesResponse struct {
	ContentType                          string
	FailResponse                         *shared.FailResponse
	Headers                              map[string][]string
	PaginatedResultPagePublicAPIResponse *shared.PaginatedResultPagePublicAPIResponse
	StatusCode                           int
	RawResponse                          *http.Response
}

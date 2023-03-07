package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteSnippetsPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type DeleteSnippetsQueryParams struct {
	Ids []string `queryParam:"style=form,explode=true,name=ids"`
}

type DeleteSnippetsHeaders struct {
	Authorization string  `header:"style=simple,explode=false,name=Authorization"`
	XRequestID    *string `header:"style=simple,explode=false,name=X-Request-ID"`
}

type DeleteSnippetsRequest struct {
	PathParams  DeleteSnippetsPathParams
	QueryParams DeleteSnippetsQueryParams
	Headers     DeleteSnippetsHeaders
}

type DeleteSnippetsResponse struct {
	ContentType    string
	DeleteResponse *shared.DeleteResponse
	FailResponse   *shared.FailResponse
	Headers        map[string][]string
	StatusCode     int
	RawResponse    *http.Response
}

package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteTermsPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type DeleteTermsQueryParams struct {
	Ids []int64 `queryParam:"style=form,explode=true,name=ids"`
}

type DeleteTermsHeaders struct {
	Authorization string  `header:"style=simple,explode=false,name=Authorization"`
	XRequestID    *string `header:"style=simple,explode=false,name=X-Request-ID"`
}

type DeleteTermsRequest struct {
	PathParams  DeleteTermsPathParams
	QueryParams DeleteTermsQueryParams
	Headers     DeleteTermsHeaders
}

type DeleteTermsResponse struct {
	ContentType    string
	DeleteResponse *shared.DeleteResponse
	FailResponse   *shared.FailResponse
	Headers        map[string][]string
	StatusCode     int
	RawResponse    *http.Response
}

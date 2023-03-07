package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type UpdateTermsPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type UpdateTermsHeaders struct {
	Authorization string  `header:"style=simple,explode=false,name=Authorization"`
	XRequestID    *string `header:"style=simple,explode=false,name=X-Request-ID"`
}

type UpdateTermsRequest struct {
	PathParams UpdateTermsPathParams
	Headers    UpdateTermsHeaders
	Request    shared.UpdateTermsRequest `request:"mediaType=application/json"`
}

type UpdateTermsResponse struct {
	ContentType         string
	CreateTermsResponse *shared.CreateTermsResponse
	FailResponse        *shared.FailResponse
	Headers             map[string][]string
	StatusCode          int
	RawResponse         *http.Response
}

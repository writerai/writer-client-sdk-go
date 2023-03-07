package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type AddTermsPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type AddTermsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type AddTermsRequest struct {
	PathParams AddTermsPathParams
	Headers    AddTermsHeaders
	Request    shared.CreateTermsRequest `request:"mediaType=application/json"`
}

type AddTermsResponse struct {
	ContentType         string
	CreateTermsResponse *shared.CreateTermsResponse
	FailResponse        *shared.FailResponse
	Headers             map[string][]string
	StatusCode          int
	RawResponse         *http.Response
}

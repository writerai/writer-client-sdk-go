package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ContetnCorrectPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type ContetnCorrectHeaders struct {
	Authorization string  `header:"style=simple,explode=false,name=Authorization"`
	XRequestID    *string `header:"style=simple,explode=false,name=X-Request-ID"`
}

type ContetnCorrectRequest struct {
	PathParams ContetnCorrectPathParams
	Headers    ContetnCorrectHeaders
	Request    shared.ContentRequest `request:"mediaType=application/json"`
}

type ContetnCorrectResponse struct {
	ContentType        string
	CorrectionResponse *shared.CorrectionResponse
	FailResponse       *shared.FailResponse
	Headers            map[string][]string
	StatusCode         int
	RawResponse        *http.Response
}

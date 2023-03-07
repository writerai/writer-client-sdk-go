package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ContentCheckPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type ContentCheckHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type ContentCheckRequest struct {
	PathParams ContentCheckPathParams
	Headers    ContentCheckHeaders
	Request    shared.ContentRequest `request:"mediaType=application/json"`
}

type ContentCheckResponse struct {
	ContentType      string
	FailResponse     *shared.FailResponse
	Headers          map[string][]string
	ProcessedContent *shared.ProcessedContent
	StatusCode       int
	RawResponse      *http.Response
}

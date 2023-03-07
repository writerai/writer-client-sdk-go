package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ContentDetectorAPIPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
}

type ContentDetectorAPIHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type ContentDetectorAPIRequest struct {
	PathParams ContentDetectorAPIPathParams
	Headers    ContentDetectorAPIHeaders
	Request    shared.ContentDetectorRequest `request:"mediaType=application/json"`
}

type ContentDetectorAPIResponse struct {
	ContentDetectorResponses []shared.ContentDetectorResponse
	ContentType              string
	FailResponse             *shared.FailResponse
	Headers                  map[string][]string
	StatusCode               int
	RawResponse              *http.Response
}

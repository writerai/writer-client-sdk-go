package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type CompletionsPathParams struct {
	ModelID        string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type CompletionsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type CompletionsRequest struct {
	PathParams CompletionsPathParams
	Headers    CompletionsHeaders
	Request    shared.CompletionRequest `request:"mediaType=application/json"`
}

type CompletionsResponse struct {
	CompletionResponse *shared.CompletionResponse
	ContentType        string
	FailResponse       *shared.FailResponse
	Headers            map[string][]string
	StatusCode         int
	RawResponse        *http.Response
}

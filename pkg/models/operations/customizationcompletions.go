package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type CustomizationCompletionsPathParams struct {
	CustomizationID string `pathParam:"style=simple,explode=false,name=customizationId"`
	ModelID         string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID  int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type CustomizationCompletionsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type CustomizationCompletionsRequest struct {
	PathParams CustomizationCompletionsPathParams
	Headers    CustomizationCompletionsHeaders
	Request    shared.CompletionRequest `request:"mediaType=application/json"`
}

type CustomizationCompletionsResponse struct {
	CompletionResponse *shared.CompletionResponse
	ContentType        string
	FailResponse       *shared.FailResponse
	Headers            map[string][]string
	StatusCode         int
	RawResponse        *http.Response
}

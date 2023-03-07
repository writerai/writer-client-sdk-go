package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ModelsPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
}

type ModelsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type ModelsRequest struct {
	PathParams ModelsPathParams
	Headers    ModelsHeaders
}

type ModelsResponse struct {
	ContentType              string
	FailResponse             *shared.FailResponse
	GenerationModelsResponse *shared.GenerationModelsResponse
	Headers                  map[string][]string
	StatusCode               int
	RawResponse              *http.Response
}

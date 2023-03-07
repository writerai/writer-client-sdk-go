package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetModelCustomizationPathParams struct {
	CustomizationID string `pathParam:"style=simple,explode=false,name=customizationId"`
	ModelID         string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID  int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type GetModelCustomizationHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type GetModelCustomizationRequest struct {
	PathParams GetModelCustomizationPathParams
	Headers    GetModelCustomizationHeaders
}

type GetModelCustomizationResponse struct {
	ContentType        string
	FailResponse       *shared.FailResponse
	Headers            map[string][]string
	ModelCustomization *shared.ModelCustomization
	StatusCode         int
	RawResponse        *http.Response
}

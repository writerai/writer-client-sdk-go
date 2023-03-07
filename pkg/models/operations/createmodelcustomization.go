package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type CreateModelCustomizationPathParams struct {
	ModelID        string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type CreateModelCustomizationHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type CreateModelCustomizationRequest struct {
	PathParams CreateModelCustomizationPathParams
	Headers    CreateModelCustomizationHeaders
	Request    shared.CreateCustomizationRequest `request:"mediaType=application/json"`
}

type CreateModelCustomizationResponse struct {
	ContentType        string
	FailResponse       *shared.FailResponse
	Headers            map[string][]string
	ModelCustomization *shared.ModelCustomization
	StatusCode         int
	RawResponse        *http.Response
}

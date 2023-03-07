package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteModelCustomizationPathParams struct {
	CustomizationID string `pathParam:"style=simple,explode=false,name=customizationId"`
	ModelID         string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID  int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type DeleteModelCustomizationRequest struct {
	PathParams DeleteModelCustomizationPathParams
}

type DeleteModelCustomizationResponse struct {
	ContentType                                      string
	FailResponse                                     *shared.FailResponse
	Headers                                          map[string][]string
	StatusCode                                       int
	RawResponse                                      *http.Response
	DeleteModelCustomization200ApplicationJSONObject map[string]interface{}
}

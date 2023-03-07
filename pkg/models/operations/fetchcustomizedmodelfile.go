package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type FetchCustomizedModelFilePathParams struct {
	CustomizationID string `pathParam:"style=simple,explode=false,name=customizationId"`
	ModelID         string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID  int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type FetchCustomizedModelFileHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type FetchCustomizedModelFileRequest struct {
	PathParams FetchCustomizedModelFilePathParams
	Headers    FetchCustomizedModelFileHeaders
}

type FetchCustomizedModelFileResponse struct {
	ContentType                                                   string
	FailResponse                                                  *shared.FailResponse
	Headers                                                       map[string][]string
	StatusCode                                                    int
	RawResponse                                                   *http.Response
	FetchCustomizedModelFile200ApplicationOctetStreamBinaryString []byte
}

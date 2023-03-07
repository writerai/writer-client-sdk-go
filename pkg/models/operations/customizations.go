package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type CustomizationsPathParams struct {
	ModelID        string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type CustomizationsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type CustomizationsRequest struct {
	PathParams CustomizationsPathParams
	Headers    CustomizationsHeaders
}

type CustomizationsResponse struct {
	ContentType            string
	CustomizationsResponse *shared.CustomizationsResponse
	FailResponse           *shared.FailResponse
	Headers                map[string][]string
	StatusCode             int
	RawResponse            *http.Response
}

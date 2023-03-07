package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetTemplateInputsPathParams struct {
	OrganizationID int64  `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64  `pathParam:"style=simple,explode=false,name=teamId"`
	TemplateID     string `pathParam:"style=simple,explode=false,name=templateId"`
}

type GetTemplateInputsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type GetTemplateInputsRequest struct {
	PathParams GetTemplateInputsPathParams
	Headers    GetTemplateInputsHeaders
}

type GetTemplateInputsResponse struct {
	ContentType             string
	FailResponse            *shared.FailResponse
	Headers                 map[string][]string
	StatusCode              int
	RawResponse             *http.Response
	TemplateDetailsResponse *shared.TemplateDetailsResponse
}

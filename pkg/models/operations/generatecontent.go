package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GenerateContentPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type GenerateContentHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type GenerateContentRequest struct {
	PathParams GenerateContentPathParams
	Headers    GenerateContentHeaders
	Request    shared.GenerateTemplateRequest `request:"mediaType=application/json"`
}

type GenerateContentResponse struct {
	ContentType  string
	Draft        *shared.Draft
	FailResponse *shared.FailResponse
	Headers      map[string][]string
	StatusCode   int
	RawResponse  *http.Response
}

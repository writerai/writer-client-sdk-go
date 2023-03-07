package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type UpdateSnippetsPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
	TeamID         int64 `pathParam:"style=simple,explode=false,name=teamId"`
}

type UpdateSnippetsHeaders struct {
	Authorization string  `header:"style=simple,explode=false,name=Authorization"`
	XRequestID    *string `header:"style=simple,explode=false,name=X-Request-ID"`
}

type UpdateSnippetsRequest struct {
	PathParams UpdateSnippetsPathParams
	Headers    UpdateSnippetsHeaders
	Request    []shared.SnippetUpdate `request:"mediaType=application/json"`
}

type UpdateSnippetsResponse struct {
	ContentType      string
	FailResponse     *shared.FailResponse
	Headers          map[string][]string
	SnippetWithUsers []shared.SnippetWithUser
	StatusCode       int
	RawResponse      *http.Response
}

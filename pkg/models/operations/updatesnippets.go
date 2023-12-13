// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type UpdateSnippetsRequest struct {
	TeamID         int64                  `pathParam:"style=simple,explode=false,name=teamId"`
	RequestBody    []shared.SnippetUpdate `request:"mediaType=application/json"`
	XRequestID     *string                `header:"style=simple,explode=false,name=X-Request-ID"`
	OrganizationID *int64                 `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *UpdateSnippetsRequest) GetTeamID() int64 {
	if o == nil {
		return 0
	}
	return o.TeamID
}

func (o *UpdateSnippetsRequest) GetRequestBody() []shared.SnippetUpdate {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateSnippetsRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

func (o *UpdateSnippetsRequest) GetOrganizationID() *int64 {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

type UpdateSnippetsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Classes     []shared.SnippetWithUser
}

func (o *UpdateSnippetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSnippetsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *UpdateSnippetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSnippetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSnippetsResponse) GetClasses() []shared.SnippetWithUser {
	if o == nil {
		return nil
	}
	return o.Classes
}

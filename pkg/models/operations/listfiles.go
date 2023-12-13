// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ListFilesRequest struct {
	OrganizationID *int64 `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *ListFilesRequest) GetOrganizationID() *int64 {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

type ListFilesResponse struct {
	// HTTP response content type for this operation
	ContentType        string
	Headers            map[string][]string
	ModelFilesResponse *shared.ModelFilesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFilesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *ListFilesResponse) GetModelFilesResponse() *shared.ModelFilesResponse {
	if o == nil {
		return nil
	}
	return o.ModelFilesResponse
}

func (o *ListFilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

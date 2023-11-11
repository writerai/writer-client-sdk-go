// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteTermsRequest struct {
	TeamID         int64   `pathParam:"style=simple,explode=false,name=teamId"`
	XRequestID     *string `header:"style=simple,explode=false,name=X-Request-ID"`
	Ids            []int64 `queryParam:"style=form,explode=true,name=ids"`
	OrganizationID *int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *DeleteTermsRequest) GetTeamID() int64 {
	if o == nil {
		return 0
	}
	return o.TeamID
}

func (o *DeleteTermsRequest) GetXRequestID() *string {
	if o == nil {
		return nil
	}
	return o.XRequestID
}

func (o *DeleteTermsRequest) GetIds() []int64 {
	if o == nil {
		return nil
	}
	return o.Ids
}

func (o *DeleteTermsRequest) GetOrganizationID() *int64 {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

type DeleteTermsResponse struct {
	// HTTP response content type for this operation
	ContentType    string
	DeleteResponse *shared.DeleteResponse
	Headers        map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteTermsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTermsResponse) GetDeleteResponse() *shared.DeleteResponse {
	if o == nil {
		return nil
	}
	return o.DeleteResponse
}

func (o *DeleteTermsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *DeleteTermsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTermsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

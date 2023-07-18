// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ListModelCustomizationsRequest struct {
	ModelID        string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID *int64 `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *ListModelCustomizationsRequest) GetModelID() string {
	if o == nil {
		return ""
	}
	return o.ModelID
}

func (o *ListModelCustomizationsRequest) GetOrganizationID() *int64 {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

type ListModelCustomizationsResponse struct {
	ContentType            string
	CustomizationsResponse *shared.CustomizationsResponse
	// Bad Request
	FailResponse *shared.FailResponse
	Headers      map[string][]string
	StatusCode   int
	RawResponse  *http.Response
}

func (o *ListModelCustomizationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListModelCustomizationsResponse) GetCustomizationsResponse() *shared.CustomizationsResponse {
	if o == nil {
		return nil
	}
	return o.CustomizationsResponse
}

func (o *ListModelCustomizationsResponse) GetFailResponse() *shared.FailResponse {
	if o == nil {
		return nil
	}
	return o.FailResponse
}

func (o *ListModelCustomizationsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListModelCustomizationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListModelCustomizationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FetchCustomizedModelFileRequest struct {
	CustomizationID string `pathParam:"style=simple,explode=false,name=customizationId"`
	ModelID         string `pathParam:"style=simple,explode=false,name=modelId"`
	OrganizationID  *int64 `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *FetchCustomizedModelFileRequest) GetCustomizationID() string {
	if o == nil {
		return ""
	}
	return o.CustomizationID
}

func (o *FetchCustomizedModelFileRequest) GetModelID() string {
	if o == nil {
		return ""
	}
	return o.ModelID
}

func (o *FetchCustomizedModelFileRequest) GetOrganizationID() *int64 {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

type FetchCustomizedModelFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse                                                   *http.Response
	FetchCustomizedModelFile200ApplicationOctetStreamBinaryString []byte
}

func (o *FetchCustomizedModelFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FetchCustomizedModelFileResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *FetchCustomizedModelFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FetchCustomizedModelFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *FetchCustomizedModelFileResponse) GetFetchCustomizedModelFile200ApplicationOctetStreamBinaryString() []byte {
	if o == nil {
		return nil
	}
	return o.FetchCustomizedModelFile200ApplicationOctetStreamBinaryString
}

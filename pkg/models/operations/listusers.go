// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type ListUsersQueryParamSortField string

const (
	ListUsersQueryParamSortFieldID               ListUsersQueryParamSortField = "id"
	ListUsersQueryParamSortFieldName             ListUsersQueryParamSortField = "name"
	ListUsersQueryParamSortFieldCreationTime     ListUsersQueryParamSortField = "creationTime"
	ListUsersQueryParamSortFieldDeleted          ListUsersQueryParamSortField = "deleted"
	ListUsersQueryParamSortFieldModificationTime ListUsersQueryParamSortField = "modificationTime"
	ListUsersQueryParamSortFieldEmail            ListUsersQueryParamSortField = "email"
	ListUsersQueryParamSortFieldLastSeen         ListUsersQueryParamSortField = "lastSeen"
)

func (e ListUsersQueryParamSortField) ToPointer() *ListUsersQueryParamSortField {
	return &e
}

func (e *ListUsersQueryParamSortField) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "name":
		fallthrough
	case "creationTime":
		fallthrough
	case "deleted":
		fallthrough
	case "modificationTime":
		fallthrough
	case "email":
		fallthrough
	case "lastSeen":
		*e = ListUsersQueryParamSortField(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUsersQueryParamSortField: %v", v)
	}
}

type ListUsersQueryParamSortOrder string

const (
	ListUsersQueryParamSortOrderAsc  ListUsersQueryParamSortOrder = "asc"
	ListUsersQueryParamSortOrderDesc ListUsersQueryParamSortOrder = "desc"
)

func (e ListUsersQueryParamSortOrder) ToPointer() *ListUsersQueryParamSortOrder {
	return &e
}

func (e *ListUsersQueryParamSortOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListUsersQueryParamSortOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUsersQueryParamSortOrder: %v", v)
	}
}

type ListUsersRequest struct {
	Limit     *int64                        `queryParam:"style=form,explode=true,name=limit"`
	Offset    *int64                        `queryParam:"style=form,explode=true,name=offset"`
	Search    *string                       `queryParam:"style=form,explode=true,name=search"`
	SortField *ListUsersQueryParamSortField `queryParam:"style=form,explode=true,name=sortField"`
	SortOrder *ListUsersQueryParamSortOrder `queryParam:"style=form,explode=true,name=sortOrder"`
}

func (o *ListUsersRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUsersRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUsersRequest) GetSearch() *string {
	if o == nil {
		return nil
	}
	return o.Search
}

func (o *ListUsersRequest) GetSortField() *ListUsersQueryParamSortField {
	if o == nil {
		return nil
	}
	return o.SortField
}

func (o *ListUsersRequest) GetSortOrder() *ListUsersQueryParamSortOrder {
	if o == nil {
		return nil
	}
	return o.SortOrder
}

type ListUsersResponse struct {
	// HTTP response content type for this operation
	ContentType                       string
	Headers                           map[string][]string
	PaginatedResultUserPublicResponse *shared.PaginatedResultUserPublicResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUsersResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ListUsersResponse) GetPaginatedResultUserPublicResponse() *shared.PaginatedResultUserPublicResponse {
	if o == nil {
		return nil
	}
	return o.PaginatedResultUserPublicResponse
}

func (o *ListUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

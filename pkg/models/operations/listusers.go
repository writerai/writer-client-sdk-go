// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

// ListUsersSortField
type ListUsersSortField string

const (
	ListUsersSortFieldID               ListUsersSortField = "id"
	ListUsersSortFieldName             ListUsersSortField = "name"
	ListUsersSortFieldCreationTime     ListUsersSortField = "creationTime"
	ListUsersSortFieldDeleted          ListUsersSortField = "deleted"
	ListUsersSortFieldModificationTime ListUsersSortField = "modificationTime"
	ListUsersSortFieldEmail            ListUsersSortField = "email"
	ListUsersSortFieldLastSeen         ListUsersSortField = "lastSeen"
)

func (e ListUsersSortField) ToPointer() *ListUsersSortField {
	return &e
}

func (e *ListUsersSortField) UnmarshalJSON(data []byte) error {
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
		*e = ListUsersSortField(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUsersSortField: %v", v)
	}
}

// ListUsersSortOrder
type ListUsersSortOrder string

const (
	ListUsersSortOrderAsc  ListUsersSortOrder = "asc"
	ListUsersSortOrderDesc ListUsersSortOrder = "desc"
)

func (e ListUsersSortOrder) ToPointer() *ListUsersSortOrder {
	return &e
}

func (e *ListUsersSortOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListUsersSortOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUsersSortOrder: %v", v)
	}
}

type ListUsersRequest struct {
	Limit     *int64              `queryParam:"style=form,explode=true,name=limit"`
	Offset    *int64              `queryParam:"style=form,explode=true,name=offset"`
	Search    *string             `queryParam:"style=form,explode=true,name=search"`
	SortField *ListUsersSortField `queryParam:"style=form,explode=true,name=sortField"`
	SortOrder *ListUsersSortOrder `queryParam:"style=form,explode=true,name=sortOrder"`
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

func (o *ListUsersRequest) GetSortField() *ListUsersSortField {
	if o == nil {
		return nil
	}
	return o.SortField
}

func (o *ListUsersRequest) GetSortOrder() *ListUsersSortOrder {
	if o == nil {
		return nil
	}
	return o.SortOrder
}

type ListUsersResponse struct {
	ContentType string
	// Bad Request
	FailResponse                      *shared.FailResponse
	Headers                           map[string][]string
	PaginatedResultUserPublicResponse *shared.PaginatedResultUserPublicResponse
	StatusCode                        int
	RawResponse                       *http.Response
}

func (o *ListUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUsersResponse) GetFailResponse() *shared.FailResponse {
	if o == nil {
		return nil
	}
	return o.FailResponse
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

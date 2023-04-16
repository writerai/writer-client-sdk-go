// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

// ListUsersSortFieldEnum
type ListUsersSortFieldEnum string

const (
	ListUsersSortFieldEnumID               ListUsersSortFieldEnum = "id"
	ListUsersSortFieldEnumName             ListUsersSortFieldEnum = "name"
	ListUsersSortFieldEnumCreationTime     ListUsersSortFieldEnum = "creationTime"
	ListUsersSortFieldEnumDeleted          ListUsersSortFieldEnum = "deleted"
	ListUsersSortFieldEnumModificationTime ListUsersSortFieldEnum = "modificationTime"
	ListUsersSortFieldEnumEmail            ListUsersSortFieldEnum = "email"
	ListUsersSortFieldEnumLastSeen         ListUsersSortFieldEnum = "lastSeen"
)

func (e *ListUsersSortFieldEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
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
		*e = ListUsersSortFieldEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUsersSortFieldEnum: %s", s)
	}
}

// ListUsersSortOrderEnum
type ListUsersSortOrderEnum string

const (
	ListUsersSortOrderEnumAsc  ListUsersSortOrderEnum = "asc"
	ListUsersSortOrderEnumDesc ListUsersSortOrderEnum = "desc"
)

func (e *ListUsersSortOrderEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "asc":
		fallthrough
	case "desc":
		*e = ListUsersSortOrderEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ListUsersSortOrderEnum: %s", s)
	}
}

type ListUsersRequest struct {
	Limit     *int64                  `queryParam:"style=form,explode=true,name=limit"`
	Offset    *int64                  `queryParam:"style=form,explode=true,name=offset"`
	Search    *string                 `queryParam:"style=form,explode=true,name=search"`
	SortField *ListUsersSortFieldEnum `queryParam:"style=form,explode=true,name=sortField"`
	SortOrder *ListUsersSortOrderEnum `queryParam:"style=form,explode=true,name=sortOrder"`
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

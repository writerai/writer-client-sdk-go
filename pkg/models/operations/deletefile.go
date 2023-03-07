package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteFilePathParams struct {
	FileID         string `pathParam:"style=simple,explode=false,name=fileId"`
	OrganizationID int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type DeleteFileRequest struct {
	PathParams DeleteFilePathParams
}

type DeleteFileResponse struct {
	ContentType                        string
	FailResponse                       *shared.FailResponse
	Headers                            map[string][]string
	StatusCode                         int
	RawResponse                        *http.Response
	DeleteFile200ApplicationJSONObject map[string]interface{}
}

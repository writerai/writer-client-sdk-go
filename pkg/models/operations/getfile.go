package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetFilePathParams struct {
	FileID         string `pathParam:"style=simple,explode=false,name=fileId"`
	OrganizationID int64  `pathParam:"style=simple,explode=false,name=organizationId"`
}

type GetFileHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type GetFileRequest struct {
	PathParams GetFilePathParams
	Headers    GetFileHeaders
}

type GetFileResponse struct {
	ContentType  string
	FailResponse *shared.FailResponse
	Headers      map[string][]string
	ModelFile    *shared.ModelFile
	StatusCode   int
	RawResponse  *http.Response
}

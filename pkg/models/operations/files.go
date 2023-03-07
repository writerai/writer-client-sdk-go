package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type FilesPathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
}

type FilesHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type FilesRequest struct {
	PathParams FilesPathParams
	Headers    FilesHeaders
}

type FilesResponse struct {
	ContentType        string
	FailResponse       *shared.FailResponse
	Headers            map[string][]string
	ModelFilesResponse *shared.ModelFilesResponse
	StatusCode         int
	RawResponse        *http.Response
}

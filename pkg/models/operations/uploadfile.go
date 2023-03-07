package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type UploadFilePathParams struct {
	OrganizationID int64 `pathParam:"style=simple,explode=false,name=organizationId"`
}

type UploadFileHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type UploadFileRequest struct {
	PathParams UploadFilePathParams
	Headers    UploadFileHeaders
	Request    shared.UploadModelFileRequest `request:"mediaType=multipart/form-data"`
}

type UploadFileResponse struct {
	ContentType  string
	FailResponse *shared.FailResponse
	Headers      map[string][]string
	ModelFile    *shared.ModelFile
	StatusCode   int
	RawResponse  *http.Response
}

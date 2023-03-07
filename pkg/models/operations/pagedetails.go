package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type PageDetailsPathParams struct {
	PageID int64 `pathParam:"style=simple,explode=false,name=pageId"`
}

type PageDetailsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type PageDetailsRequest struct {
	PathParams PageDetailsPathParams
	Headers    PageDetailsHeaders
}

type PageDetailsResponse struct {
	ContentType             string
	FailResponse            *shared.FailResponse
	Headers                 map[string][]string
	PageWithSectionResponse *shared.PageWithSectionResponse
	StatusCode              int
	RawResponse             *http.Response
}

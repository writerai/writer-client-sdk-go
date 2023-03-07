package operations

import (
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetSubscriptionDetailsHeaders struct {
	Authorization string `header:"style=simple,explode=false,name=Authorization"`
}

type GetSubscriptionDetailsRequest struct {
	Headers GetSubscriptionDetailsHeaders
}

type GetSubscriptionDetailsResponse struct {
	ContentType                   string
	FailResponse                  *shared.FailResponse
	Headers                       map[string][]string
	StatusCode                    int
	RawResponse                   *http.Response
	SubscriptionPublicResponseAPI *shared.SubscriptionPublicResponseAPI
}

<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithSecurity(""),
		writerclientsdkgo.WithOrganizationID(496531),
	)

	contentDetectorRequest := shared.ContentDetectorRequest{
		Input: "Bronze Indian",
	}

	var organizationID *int64 = 558689

	ctx := context.Background()
	res, err := s.AIContentDetector.Detect(ctx, contentDetectorRequest, organizationID)
	if err != nil {
		log.Fatal(err)
	}

	if res.ContentDetectorResponses != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->
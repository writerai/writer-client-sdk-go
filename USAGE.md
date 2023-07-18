<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(548814),
    )

    ctx := context.Background()
    res, err := s.AIContentDetector.Detect(ctx, operations.DetectContentRequest{
        ContentDetectorRequest: shared.ContentDetectorRequest{
            Input: "provident",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContentDetectorResponses != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/writerai/writer-client-sdk-go"
    "github.com/writerai/writer-client-sdk-go/pkg/models/shared"
    "github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writerai.New()
    
    req := operations.ContentDetectorAPIRequest{
        PathParams: operations.ContentDetectorAPIPathParams{
            OrganizationID: 548814,
        },
        Headers: operations.ContentDetectorAPIHeaders{
            Authorization: "deserunt",
        },
        Request: shared.ContentDetectorRequest{
            Input: "porro",
        },
    }

    ctx := context.Background()
    res, err := s.AIContentDetector.ContentDetectorAPI(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ContentDetectorResponses != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->
# DownloadTheCustomizedModel

## Overview

Methods related to Download the customized model

### Available Operations

* [FetchFile](#fetchfile) - Download your fine-tuned model (available only for Palmyra Base and Palmyra Large)

## FetchFile

Download your fine-tuned model (available only for Palmyra Base and Palmyra Large)

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(681820),
    )

    ctx := context.Background()
    res, err := s.DownloadTheCustomizedModel.FetchFile(ctx, operations.FetchCustomizedModelFileRequest{
        CustomizationID: "in",
        ModelID: "corporis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchCustomizedModelFile200ApplicationOctetStreamBinaryString != nil {
        // handle response
    }
}
```

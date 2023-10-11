# DownloadTheCustomizedModel
(*DownloadTheCustomizedModel*)

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
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(501762),
    )

    ctx := context.Background()
    res, err := s.DownloadTheCustomizedModel.FetchFile(ctx, operations.FetchCustomizedModelFileRequest{
        CustomizationID: "Racine beyond connecting",
        ModelID: "invoice Folk",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchCustomizedModelFile200ApplicationOctetStreamBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.FetchCustomizedModelFileRequest](../../models/operations/fetchcustomizedmodelfilerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.FetchCustomizedModelFileResponse](../../models/operations/fetchcustomizedmodelfileresponse.md), error**


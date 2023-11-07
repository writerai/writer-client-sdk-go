# Files
(*.Files*)

## Overview

Methods related to Files

### Available Operations

* [Delete](#delete) - Delete file
* [Get](#get) - Get file
* [List](#list) - List files
* [Upload](#upload) - Upload file

## Delete

Delete file

### Example Usage

```go
package main

import(
	"context"
	"log"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(545907),
    )


    var fileID string = "string"

    var organizationID *int64 = 841399

    ctx := context.Background()
    res, err := s.Files.Delete(ctx, fileID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `fileID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.DeleteFileResponse](../../models/operations/deletefileresponse.md), error**


## Get

Get file

### Example Usage

```go
package main

import(
	"context"
	"log"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(700347),
    )


    var fileID string = "string"

    var organizationID *int64 = 90065

    ctx := context.Background()
    res, err := s.Files.Get(ctx, fileID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `fileID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.GetFileResponse](../../models/operations/getfileresponse.md), error**


## List

List files

### Example Usage

```go
package main

import(
	"context"
	"log"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(768578),
    )


    var organizationID *int64 = 99895

    ctx := context.Background()
    res, err := s.Files.List(ctx, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.ListFilesResponse](../../models/operations/listfilesresponse.md), error**


## Upload

Upload file

### Example Usage

```go
package main

import(
	"context"
	"log"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity(""),
        writerclientsdkgo.WithOrganizationID(403654),
    )


    uploadModelFileRequest := shared.UploadModelFileRequest{
        File: shared.File{
            Content: []byte("0x7cbca97eC6"),
            FileName: "plastic_cli.gif",
        },
    }

    var organizationID *int64 = 360896

    ctx := context.Background()
    res, err := s.Files.Upload(ctx, uploadModelFileRequest, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `uploadModelFileRequest`                                                       | [shared.UploadModelFileRequest](../../models/shared/uploadmodelfilerequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `organizationID`                                                               | **int64*                                                                       | :heavy_minus_sign:                                                             | N/A                                                                            |


### Response

**[*operations.UploadFileResponse](../../models/operations/uploadfileresponse.md), error**


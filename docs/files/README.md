# Files

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
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(613064),
    )

    ctx := context.Background()
    res, err := s.Files.Delete(ctx, operations.DeleteFileRequest{
        FileID: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteFile200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Get

Get file

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
        writer.WithOrganizationID(902349),
    )

    ctx := context.Background()
    res, err := s.Files.Get(ctx, operations.GetFileRequest{
        FileID: "quidem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelFile != nil {
        // handle response
    }
}
```

## List

List files

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
        writer.WithOrganizationID(99280),
    )

    ctx := context.Background()
    res, err := s.Files.List(ctx, operations.ListFilesRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelFilesResponse != nil {
        // handle response
    }
}
```

## Upload

Upload file

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(60225),
    )

    ctx := context.Background()
    res, err := s.Files.Upload(ctx, operations.UploadFileRequest{
        UploadModelFileRequest: shared.UploadModelFileRequest{
            File: shared.UploadModelFileRequestFile{
                Content: []byte("reiciendis"),
                File: "est",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelFile != nil {
        // handle response
    }
}
```
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
    req := operations.DeleteFileRequest{
        FileID: "iure",
    }

    res, err := s.Files.Delete(ctx, req)
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
    req := operations.GetFileRequest{
        FileID: "quidem",
    }

    res, err := s.Files.Get(ctx, req)
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
    req := operations.ListFilesRequest{}

    res, err := s.Files.List(ctx, req)
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
    req := operations.UploadFileRequest{
        UploadModelFileRequest: shared.UploadModelFileRequest{
            File: shared.UploadModelFileRequestFile{
                Content: []byte("reiciendis"),
                File: "est",
            },
        },
    }

    res, err := s.Files.Upload(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelFile != nil {
        // handle response
    }
}
```

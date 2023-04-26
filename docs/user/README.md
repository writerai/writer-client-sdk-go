# User

## Overview

Methods related to User

### Available Operations

* [List](#list) - List users

## List

List users

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
        writer.WithOrganizationID(55),
    )

    ctx := context.Background()    
    req := operations.ListUsersRequest{
        Limit: writer.Int64(872651),
        Offset: writer.Int64(311860),
        Search: writer.String("tempora"),
        SortField: operations.ListUsersSortFieldEnumCreationTime.ToPointer(),
        SortOrder: operations.ListUsersSortOrderEnumDesc.ToPointer(),
    }

    res, err := s.User.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResultUserPublicResponse != nil {
        // handle response
    }
}
```

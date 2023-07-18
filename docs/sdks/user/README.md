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
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "",
        }),
        writer.WithOrganizationID(55),
    )

    ctx := context.Background()
    res, err := s.User.List(ctx, operations.ListUsersRequest{
        Limit: writer.Int64(872651),
        Offset: writer.Int64(311860),
        Search: writer.String("tempora"),
        SortField: operations.ListUsersSortFieldCreationTime.ToPointer(),
        SortOrder: operations.ListUsersSortOrderDesc.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResultUserPublicResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListUsersRequest](../../models/operations/listusersrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**


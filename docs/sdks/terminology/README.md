# Terminology
(*Terminology*)

## Overview

Methods related to Terminology

### Available Operations

* [Add](#add) - Add terms
* [Delete](#delete) - Delete terms
* [Find](#find) - Find terms
* [Update](#update) - Update terms

## Add

Add terms

### Example Usage

```go
package main

import(
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(551477),
    )


    createTermsRequest := shared.CreateTermsRequest{}

    var teamID int64 = 823436

    var organizationID *int64 = writerclientsdkgo.Int64(554561)

    ctx := context.Background()
    res, err := s.Terminology.Add(ctx, createTermsRequest, teamID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `createTermsRequest`                                                       | [shared.CreateTermsRequest](../../pkg/models/shared/createtermsrequest.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `teamID`                                                                   | *int64*                                                                    | :heavy_check_mark:                                                         | N/A                                                                        |
| `organizationID`                                                           | **int64*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |


### Response

**[*operations.AddTermsResponse](../../pkg/models/operations/addtermsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## Delete

Delete terms

### Example Usage

```go
package main

import(
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(545907),
    )


    var teamID int64 = 841399

    var xRequestID *string = writerclientsdkgo.String("<value>")

    ids := []int64{
        698486,
    }

    var organizationID *int64 = writerclientsdkgo.Int64(557937)

    ctx := context.Background()
    res, err := s.Terminology.Delete(ctx, teamID, xRequestID, ids, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `teamID`                                              | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `xRequestID`                                          | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `ids`                                                 | []*int64*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.DeleteTermsResponse](../../pkg/models/operations/deletetermsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## Find

Find terms

### Example Usage

```go
package main

import(
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(40141),
    )

    ctx := context.Background()
    res, err := s.Terminology.Find(ctx, operations.FindTermsRequest{
        TeamID: 326883,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedResultFullTermWithUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.FindTermsRequest](../../pkg/models/operations/findtermsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.FindTermsResponse](../../pkg/models/operations/findtermsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

## Update

Update terms

### Example Usage

```go
package main

import(
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"context"
	"log"
)

func main() {
    s := writerclientsdkgo.New(
        writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        writerclientsdkgo.WithOrganizationID(857478),
    )


    updateTermsRequest := shared.UpdateTermsRequest{}

    var teamID int64 = 24555

    var xRequestID *string = writerclientsdkgo.String("<value>")

    var organizationID *int64 = writerclientsdkgo.Int64(597129)

    ctx := context.Background()
    res, err := s.Terminology.Update(ctx, updateTermsRequest, teamID, xRequestID, organizationID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `updateTermsRequest`                                                       | [shared.UpdateTermsRequest](../../pkg/models/shared/updatetermsrequest.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `teamID`                                                                   | *int64*                                                                    | :heavy_check_mark:                                                         | N/A                                                                        |
| `xRequestID`                                                               | **string*                                                                  | :heavy_minus_sign:                                                         | N/A                                                                        |
| `organizationID`                                                           | **int64*                                                                   | :heavy_minus_sign:                                                         | N/A                                                                        |


### Response

**[*operations.UpdateTermsResponse](../../pkg/models/operations/updatetermsresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 4xx-5xx                | */*                    |

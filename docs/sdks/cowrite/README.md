# CoWrite
(*CoWrite*)

## Overview

Methods related to CoWrite

### Available Operations

* [GenerateContent](#generatecontent) - Generate content using predefined templates
* [ListTemplates](#listtemplates) - Get a list of your existing CoWrite templates

## GenerateContent

Generate content using predefined templates

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
        writerclientsdkgo.WithOrganizationID(569932),
    )


    generateTemplateRequest := shared.GenerateTemplateRequest{
        Inputs: []shared.MagicRequestInput{
            shared.MagicRequestInput{
                Name: "string",
                Value: []string{
                    "string",
                },
            },
        },
        TemplateID: "string",
    }

    var teamID int64 = 888452

    var organizationID *int64 = 926220

    ctx := context.Background()
    res, err := s.CoWrite.GenerateContent(ctx, generateTemplateRequest, teamID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Draft != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `generateTemplateRequest`                                                               | [shared.GenerateTemplateRequest](../../../pkg/models/shared/generatetemplaterequest.md) | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `teamID`                                                                                | *int64*                                                                                 | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `organizationID`                                                                        | **int64*                                                                                | :heavy_minus_sign:                                                                      | N/A                                                                                     |


### Response

**[*operations.GenerateContentResponse](../../pkg/models/operations/generatecontentresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 400-600                | */*                    |

## ListTemplates

Get a list of your existing CoWrite templates

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
        writerclientsdkgo.WithOrganizationID(380445),
    )


    var teamID int64 = 882866

    var templateID string = "string"

    var organizationID *int64 = 55511

    ctx := context.Background()
    res, err := s.CoWrite.ListTemplates(ctx, teamID, templateID, organizationID)
    if err != nil {
        log.Fatal(err)
    }

    if res.TemplateDetailsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `teamID`                                              | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `templateID`                                          | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `organizationID`                                      | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.ListTemplatesResponse](../../pkg/models/operations/listtemplatesresponse.md), error**
| Error Object           | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.FailResponse | 400,401,403,404,500    | application/json       |
| sdkerrors.SDKError     | 400-600                | */*                    |

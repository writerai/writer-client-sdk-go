# Content

## Overview

Methods related to Content

### Available Operations

* [Check](#check) - Check your content against your preset styleguide.
* [Correct](#correct) - Apply the style guide suggestions directly to your content.

## Check

Check your content against your preset styleguide.

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
        writer.WithOrganizationID(832620),
    )

    ctx := context.Background()
    res, err := s.Content.Check(ctx, operations.ContentCheckRequest{
        ContentRequest: shared.ContentRequest{
            Content: "sapiente",
            Settings: shared.ContentSettings{
                AgeAndFamilyStatus: false,
                Confidence: false,
                ContentSafeguards: false,
                Disability: false,
                GenderIdentitySensitivity: false,
                GenderInclusiveNouns: false,
                GenderInclusivePronouns: false,
                Grammar: false,
                HealthyCommunication: false,
                PassiveVoice: false,
                RaceEthnicityNationalitySensitivity: false,
                SexualOrientationSensitivity: false,
                Spelling: false,
                SubstanceUseSensitivity: false,
                UnclearReference: false,
                Wordiness: false,
            },
        },
        TeamID: 778157,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProcessedContent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ContentCheckRequest](../../models/operations/contentcheckrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ContentCheckResponse](../../models/operations/contentcheckresponse.md), error**


## Correct

Apply the style guide suggestions directly to your content.

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
        writer.WithOrganizationID(140350),
    )

    ctx := context.Background()
    res, err := s.Content.Correct(ctx, operations.ContentCorrectRequest{
        ContentRequest: shared.ContentRequest{
            Content: "at",
            Settings: shared.ContentSettings{
                AgeAndFamilyStatus: false,
                Confidence: false,
                ContentSafeguards: false,
                Disability: false,
                GenderIdentitySensitivity: false,
                GenderInclusiveNouns: false,
                GenderInclusivePronouns: false,
                Grammar: false,
                HealthyCommunication: false,
                PassiveVoice: false,
                RaceEthnicityNationalitySensitivity: false,
                SexualOrientationSensitivity: false,
                Spelling: false,
                SubstanceUseSensitivity: false,
                UnclearReference: false,
                Wordiness: false,
            },
        },
        XRequestID: writer.String("at"),
        TeamID: 978619,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CorrectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ContentCorrectRequest](../../models/operations/contentcorrectrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ContentCorrectResponse](../../models/operations/contentcorrectresponse.md), error**


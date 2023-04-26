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
	"github.com/writerai/writer-client-sdk-go/pkg/models/operations"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
)

func main() {
    s := writer.New(
        writer.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
        writer.WithOrganizationID(149675),
    )

    ctx := context.Background()    
    req := operations.ContentCheckRequest{
        ContentRequest: shared.ContentRequest{
            Content: "iste",
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
        TeamID: 222321,
    }

    res, err := s.Content.Check(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ProcessedContent != nil {
        // handle response
    }
}
```

## Correct

Apply the style guide suggestions directly to your content.

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
        writer.WithOrganizationID(616934),
    )

    ctx := context.Background()    
    req := operations.ContentCorrectRequest{
        ContentRequest: shared.ContentRequest{
            Content: "laboriosam",
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
        XRequestID: writer.String("hic"),
        TeamID: 902599,
    }

    res, err := s.Content.Correct(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CorrectionResponse != nil {
        // handle response
    }
}
```

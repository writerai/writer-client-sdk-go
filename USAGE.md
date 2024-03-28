<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	writerclientsdkgo "github.com/writerai/writer-client-sdk-go"
	"github.com/writerai/writer-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := writerclientsdkgo.New(
		writerclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
		writerclientsdkgo.WithOrganizationID(850421),
	)

	ctx := context.Background()
	res, err := s.Billing.GetSubscriptionDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.SubscriptionPublicResponseAPI != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->
# CompletionRequest


## Fields

| Field                  | Type                   | Required               | Description            | Example                |
| ---------------------- | ---------------------- | ---------------------- | ---------------------- | ---------------------- |
| `BestOf`               | **int64*               | :heavy_minus_sign:     | N/A                    | 1                      |
| `FrequencyPenalty`     | **float64*             | :heavy_minus_sign:     | N/A                    |                        |
| `Logprobs`             | **int64*               | :heavy_minus_sign:     | N/A                    |                        |
| `MaxTokens`            | **int64*               | :heavy_minus_sign:     | N/A                    | 1024                   |
| `MinTokens`            | **int64*               | :heavy_minus_sign:     | N/A                    | 1                      |
| `N`                    | **int64*               | :heavy_minus_sign:     | N/A                    |                        |
| `PresencePenalty`      | **float64*             | :heavy_minus_sign:     | N/A                    |                        |
| `Prompt`               | *string*               | :heavy_check_mark:     | N/A                    |                        |
| `Stop`                 | []*string*             | :heavy_minus_sign:     | N/A                    | [<br/>"the",<br/>"is",<br/>"and"<br/>] |
| `Temperature`          | **float64*             | :heavy_minus_sign:     | N/A                    | 0.7                    |
| `TopP`                 | **float64*             | :heavy_minus_sign:     | N/A                    | 1                      |
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContentDetectorRequest struct {
	Input string `json:"input"`
}

func (o *ContentDetectorRequest) GetInput() string {
	if o == nil {
		return ""
	}
	return o.Input
}

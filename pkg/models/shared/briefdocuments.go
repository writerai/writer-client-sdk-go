// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BriefDocuments struct {
	Pagination Pagination      `json:"pagination"`
	Result     []BriefDocument `json:"result,omitempty"`
	TotalCount int64           `json:"totalCount"`
}
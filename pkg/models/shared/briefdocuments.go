// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BriefDocuments struct {
	Pagination Pagination      `json:"pagination"`
	Result     []BriefDocument `json:"result,omitempty"`
	TotalCount int64           `json:"totalCount"`
}

func (o *BriefDocuments) GetPagination() Pagination {
	if o == nil {
		return Pagination{}
	}
	return o.Pagination
}

func (o *BriefDocuments) GetResult() []BriefDocument {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *BriefDocuments) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

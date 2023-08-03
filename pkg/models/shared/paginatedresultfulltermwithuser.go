// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaginatedResultFullTermWithUser struct {
	Pagination Pagination         `json:"pagination"`
	Result     []FullTermWithUser `json:"result,omitempty"`
	TotalCount int64              `json:"totalCount"`
}

func (o *PaginatedResultFullTermWithUser) GetPagination() Pagination {
	if o == nil {
		return Pagination{}
	}
	return o.Pagination
}

func (o *PaginatedResultFullTermWithUser) GetResult() []FullTermWithUser {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *PaginatedResultFullTermWithUser) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

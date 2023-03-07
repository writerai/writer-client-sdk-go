package shared

type PaginatedResultPagePublicAPIResponse struct {
	Pagination Pagination              `json:"pagination"`
	Result     []PagePublicAPIResponse `json:"result,omitempty"`
	TotalCount int64                   `json:"totalCount"`
}

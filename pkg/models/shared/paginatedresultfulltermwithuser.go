package shared

type PaginatedResultFullTermWithUser struct {
	Pagination Pagination         `json:"pagination"`
	Result     []FullTermWithUser `json:"result,omitempty"`
	TotalCount int64              `json:"totalCount"`
}

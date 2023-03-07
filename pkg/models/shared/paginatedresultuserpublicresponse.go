package shared

type PaginatedResultUserPublicResponse struct {
	Pagination Pagination           `json:"pagination"`
	Result     []UserPublicResponse `json:"result,omitempty"`
	TotalCount int64                `json:"totalCount"`
}

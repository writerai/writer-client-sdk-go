package shared

type PaginatedResultSnippetWithUser struct {
	Pagination Pagination        `json:"pagination"`
	Result     []SnippetWithUser `json:"result,omitempty"`
	TotalCount int64             `json:"totalCount"`
}

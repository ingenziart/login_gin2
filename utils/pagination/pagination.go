package pagination

type PaginationResponse struct {
	Status       string      `json:"status"`
	List         interface{} `json:"list"`
	Total        int64       `json:"total"`
	PreviousPage *int        `json:"previousPage"`
	NextPage     *int        `json:"nextPage"`
	LastPage     int         `json:"lastPage"`
	CurrentPage  int         `json:"currentPage"`
}

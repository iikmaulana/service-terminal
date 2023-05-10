package models

type FilterParams struct {
	Page     int64
	Limit    int64
	SortType string
	Filter   []string
	Sparator string
	Search   []string
	All      string
}

type FormMetaData struct {
	TotalDataPerpage int
	TotalPage        int
	TotalData        int64
	To               int
	From             int
}

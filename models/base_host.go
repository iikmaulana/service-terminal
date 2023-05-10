package models

type HostCreateRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Port string `json:"port"`
}

type HostUpdateRequest struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Url  string `json:"url" db:"url"`
	Port string `json:"port" db:"port"`
}

type HostResult struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Url  string `json:"url" db:"url"`
	Port string `json:"port" db:"port"`
}

type HostListResult struct {
	TotalData        int64        `json:"total_data"`
	TotalDataPerpage int64        `json:"total_data_perpage"`
	From             int64        `json:"from"`
	To               int64        `json:"to"`
	TotalPage        int64        `json:"total_page"`
	Data             []HostResult `json:"data"`
}

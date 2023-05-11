package models

type HostCreateRequest struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	Port         string `json:"port"`
	HostUsername string `json:"host_username"`
	HostPassword string `json:"host_password"`
	HostClientId string `json:"host_client_id"`
	HostType     string `json:"host_type"`
}

type HostUpdateRequest struct {
	Id           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Url          string `json:"url" db:"url"`
	Port         string `json:"port" db:"port"`
	HostUsername string `json:"host_username" db:"host_username"`
	HostPassword string `json:"host_password" db:"host_password"`
	HostClientId string `json:"host_client_id" db:"host_client_id"`
	HostType     string `json:"host_type" db:"host_type"`
}

type HostResult struct {
	Id           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Url          string `json:"url" db:"url"`
	Port         string `json:"port" db:"port"`
	HostUsername string `json:"host_username" db:"host_username"`
	HostPassword string `json:"host_password" db:"host_password"`
	HostClientId string `json:"host_client_id" db:"host_client_id"`
	HostType     string `json:"host_type" db:"host_type"`
}

type HostListResult struct {
	TotalData        int64        `json:"total_data"`
	TotalDataPerpage int64        `json:"total_data_perpage"`
	From             int64        `json:"from"`
	To               int64        `json:"to"`
	TotalPage        int64        `json:"total_page"`
	Data             []HostResult `json:"data"`
}

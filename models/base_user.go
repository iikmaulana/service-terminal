package models

type UserCreateRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	Status    int    `json:"status"`
	LastLogin string `json:"last_login"`
}

type UserUpdateRequest struct {
	Id        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Status    int    `json:"status" db:"status"`
	LastLogin string `json:"last_login" db:"last_login"`
}

type UserResult struct {
	Id        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Status    int    `json:"status" db:"status"`
	LastLogin string `json:"last_login" db:"last_login"`
}

type UserViewResult struct {
	Id        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Status    int    `json:"status" db:"status"`
	LastLogin string `json:"last_login" db:"last_login"`
}

type UserListResult struct {
	TotalData        int64            `json:"total_data"`
	TotalDataPerpage int64            `json:"total_data_perpage"`
	From             int64            `json:"from"`
	To               int64            `json:"to"`
	TotalPage        int64            `json:"total_page"`
	Data             []UserViewResult `json:"data"`
}

package query

const (
	QueryUserInsert = `INSERT INTO db_terminal.user (id,username,password,created_at,status,last_login) VALUES ($1,$2,$3,$4,$5,$6) returning id`
	QueryUserList   = `SELECT id,username,created_at,status,last_login FROM db_terminal.user`
	QueryUserView   = `SELECT id,username,password,created_at,status,last_login FROM db_terminal.user WHERE id = $1`
	QueryUserDelete = `DELETE FROM db_terminal.user WHERE id = $1`
	QueryUserLogin  = `SELECT id,username,password,created_at,status,last_login FROM db_terminal.user WHERE username = $1`
)

package query

const (
	QueryHostInsert = `INSERT INTO db_terminal.host (id,name,url,port) VALUES ($1,$2,$3,$4) returning id`
	QueryHostList   = `SELECT id,name,url,port FROM db_terminal.host`
	QueryHostView   = `SELECT id,name,url,port FROM db_terminal.host WHERE id = $1`
	QueryHostDelete = `DELETE FROM db_terminal.host WHERE id = $1`
)

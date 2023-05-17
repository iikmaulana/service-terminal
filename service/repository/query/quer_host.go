package query

const (
	QueryHostInsert = `INSERT INTO db_terminal.host (id,name,url,port,host_username,host_password,host_client_id,host_type,topic) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) returning id`
	QueryHostList   = `SELECT id,name,url,port,host_username,host_password,host_client_id,host_type,topic FROM db_terminal.host`
	QueryHostView   = `SELECT id,name,url,port,host_username,host_password,host_client_id,host_type,topic FROM db_terminal.host WHERE id = $1`
	QueryHostDelete = `DELETE FROM db_terminal.host WHERE id = $1`
)

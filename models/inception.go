package models

import 	(
	mysql "database/sql"
	)
type Inception struct {
	Stage         string `json:"stage"`
	Order_id      string `json:"order_id"`
	Error_level   string `json:"error_level"`
	Stage_status  string `json:"stage_status"`
	Error_message mysql.NullString `json:"error_message"`
	Sql           string `json:"sql"`
	Affected_rows string `json:"affected_rows"`
	Sequence      string `json:"sequence"`
	Backup_dbname mysql.NullString `json:"backup_dbname"`
	Execute_time  mysql.NullString `json:"execute_time"`
	Sqlsha1       mysql.NullString `json:"sqlsha1"`
	Backup_time   string `json:"backup_time"`
}
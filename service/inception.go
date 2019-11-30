package service

import (
	mysql "database/sql"
	"fmt"
	"rds/goinception/v1/conf"

	//"rds/goinception/v1/handler"
	"rds/goinception/v1/infra"
	"rds/goinception/v1/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type CreateInception struct {
	SqlStr   string `json:"sql_str"`
	SqlTypes int    `json:"sql_types"`
	User     string `json:"user"`
	Pwd      string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type ClusterInfo struct {
	User string `json:"user"`
	Pwd  string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func CheckHead(info ClusterInfo) (sqlHead string) {
	checkHead := fmt.Sprintf(conf.CheckHead, info.User, info.Pwd, info.Host, info.Port)
	return checkHead
}
func ExecuteHead(info ClusterInfo) (sqlHead string) {
	executeHead := fmt.Sprintf(conf.ExecuteHead, info.User, info.Pwd, info.Host, info.Port)
	return executeHead
}
func InceptionBody(info ClusterInfo, database string, sqlText string, level int) (incHead string, err error) {
	// level 预审还是执行， 1预审 2执行
	sqlBody := fmt.Sprintf(conf.InceptionBody, database, sqlText)
	head := ""
	if level == 1 {
		head = CheckHead(info)
	} else if level == 2 {
		head = ExecuteHead(info)
	}
	incHead = head + sqlBody
	return incHead, nil
}

func InceptionTest(createInc CreateInception) (ml []models.Inception, err error) {
	// 连接inception 服务器
	db, err := mysql.Open("mysql",
		infra.Conf.Inception.Url)
	if err != nil {
		fmt.Println("err: ", err)
	}
	var info = ClusterInfo{
		User: createInc.User,
		Pwd:  createInc.Pwd,
		Host: createInc.Host,
		Port: createInc.Port,
	}
	sqlStr, err := InceptionBody(info, createInc.Database, createInc.SqlStr, createInc.SqlTypes)
	var inc models.Inception
	rows, err := db.Query(sqlStr)
	if err != nil {
		logrus.Error(err.Error())
		return nil, errors.Wrap(err, "查询数据库出差")
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&inc.Stage, &inc.Order_id, &inc.Error_level, &inc.Stage_status, &inc.Error_message, &inc.Sql, &inc.Affected_rows,
			&inc.Sequence, &inc.Backup_dbname, &inc.Execute_time, &inc.Sqlsha1, &inc.Backup_time)
		//fmt.Println(err, inc)
		//logrus.Info("err", err, "result:", inc)
		ml = append(ml, inc)
	}
	for _, v := range ml {
		if v.Error_level == "2" {
			logrus.Error("err", v)
		}
	}

	return ml, nil
}

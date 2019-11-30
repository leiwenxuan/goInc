package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	sql_test := "CREATE TABLE `t_rds_cluster_%s` (" +
		" `id` int(11) NOT NULL AUTO_INCREMENT," +
		" `status` int(11) DEFAULT NULL," +
		"`cluster_id` varchar(128) DEFAULT NULL," +
		"  `cluster_name` varchar(128) NOT NULL, " +
		"  `cluster_type` varchar(64) DEFAULT NULL," +
		"`create_time` datetime(6) NOT NULL," +
		"`current_job` varchar(128) DEFAULT NULL," +
		"`spec_id` int(11) NOT NULL COMMENT '规格id'," +
		"`spec_name` varchar(128) NOT NULL," +
		"`team_work` varchar(128) NOT NULL," +
		"`cluster_note` varchar(64) NOT NULL," +
		"`cluster_mark` varchar(128) DEFAULT NULL," +
		"`create_name` varchar(32) NOT NULL," +
		"`source` int(11) NOT NULL," +
		"PRIMARY KEY (`id`)," +
		"UNIQUE KEY `cluster_id` (`cluster_id`)" +
		"); \n"
	start := time.Now()

	f, err3 := os.Create("example/test.sql") //创建文件
	defer f.Close()
	if err3 != nil {
		logrus.Info("创建文件失败：", err3.Error())
	}
	for i := 0; i < 1000; i++ {
		str := strconv.Itoa(i)

		sql_text := fmt.Sprintf(sql_test, str)
		_, err := f.WriteString(sql_text) //写入文件(字节数组)
		if err != nil {
			logrus.Error("写入文件失败:", err.Error())
		}
		f.Sync()
	}

	cost := time.Since(start)
	fmt.Printf("执行inception 运行时间 =[%s]", cost)
	//n2, err3 := f.Write([]byte("雷文轩")) //写入文件(字节数组)
	//fmt.Printf("写入 %d 个字节n", n2)
	//n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	//fmt.Printf("写入 %d 个字节n", n3)
	//f.Sync()
}

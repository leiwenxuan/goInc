package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"rds/goinception/v1/infra"
	"rds/goinception/v1/router"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func readFile(filePath *string) (str string) {

	fi, err := os.Open(*filePath)
	if err != nil {
		logrus.Error("读取文件错误：", err.Error())
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)

}

//var info = service.ClusterInfo{
//	User: "root",
//	Pwd:  "Root@123",
//	Host: "47.107.202.130",
//	Port: "3306",
//}

func main() {
	// 加载配置文件
	infra.InitConfig()
	infra.SetLineNumLogrusHook()
	// 加载router
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()
	router.InitRouter(g)
	logrus.Info(http.ListenAndServe(":9000", g).Error())

	//dbList := []string{"test_01", "test_02", "test_03", "test_04", "test_05", "test_06", "test_07",
	//	"test_08", "test_09", "test_10", "test_11", "test_12", "test_13", "test_15"}
	//var info = service.ClusterInfo{
	//	User: "root",
	//	Pwd:  "123456",
	//	Host: "192.168.1.3",
	//	Port: "20001",
	//}
	//filePath := flag.String("file", "example/test.sql", "读取文件目录")
	//var li []string
	//for _, v := range dbList {
	//	start := time.Now()
	//	//some func or operation
	//	str := readFile(filePath)
	//	logrus.Info("读取文件： ", time.Since(start))
	//	result, err := service.InceptionTest(v, str, info)
	//	if err != nil {
	//		logrus.Info("执行inception错误")
	//	}
	//	//bsql, err := json.Marshal(result)
	//	//if err != nil {
	//	//	panic(err)
	//	//}
	//	//fmt.Println(string(bsql))
	//	cost := time.Since(start)
	//	li = append(li, fmt.Sprintf("%s", cost))
	//	fmt.Printf("执行inception 运行时间 =[%s]", cost)
	//	logrus.Info("总的个数： ", len(result))
	//}
	//fmt.Println(li)
}

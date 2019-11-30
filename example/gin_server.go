package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
)

func ginConsulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}

func ginRegisterServer(serverId string, port int, checkPort int) {
	time.Sleep(time.Second * 3)
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	//checkPort := 18080

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = serverId
	registration.Name = "serverNode"
	registration.Port = port
	registration.Tags = []string{"serverNode"}
	registration.Address = "192.168.1.3"
	fmt.Println(port)
	fmt.Println(fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"))
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s", //check失败后30秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

	http.HandleFunc("/check", ginConsulCheck)
	err = http.ListenAndServe(fmt.Sprintf(":%d", 18080), nil)
	if err != nil {
		fmt.Println("err", err.Error())
	}
}
func main() {
	r := gin.Default()
	r.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	count := 9528
	url := fmt.Sprintf(":%d", count)
	ginRegisterServer(fmt.Sprintf("%d", count), count, count)

	fmt.Println("url: ", url)
	r.Run(url) // 在 0.0.0.0:8080 上监听并服务

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	fmt.Println(client, err)
	for {

		time.Sleep(time.Second * 3)
		var services map[string]*consulapi.AgentService
		var err error

		services, err = client.Agent().Services()

		if nil != err {
			log.Println("in consual list Services:", err)
			continue
		}

		for i := 9528; i < 9532; i++ {
			url := fmt.Sprintf("%d", i)

			if _, found := services[url]; !found {
				log.Println(url, "server  not found")
				continue
			}
			sendData(services[url])
		}

	}
}

func sendData(service *consulapi.AgentService) {
	fmt.Println(service.Address, service.Port)
	//conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", service.Address, service.Port))
	//conn, err := net.Dial("tcp", "192.168.1.3:9527")
	fmt.Println(fmt.Sprintf("http://%s:%d/check", service.Address, service.Port))
	url := fmt.Sprintf("http://%s:%d/check", service.Address, service.Port)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Body)
}

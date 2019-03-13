package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"testing"
)

const Id = "1234567890"

func main() { //TestRegister(t *testing.T)

	fmt.Println("test begin .")
	config := consulapi.DefaultConfig()
	config.Address = "10.10.6.3:8500"
	fmt.Println("defautl config : ", config)
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	//创建一个新服务。
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = Id
	registration.Name = "go-web"
	registration.Port = 8502
	registration.Tags = []string{"go-web"}
	registration.Address = "192.168.21.236"

	//增加check。
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/health")
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "10s"
	//注册check服务。
	registration.Check = check
	log.Println("get check.HTTP:", check)

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

}

func TestDregister(t *testing.T) {

	fmt.Println("test begin .")
	config := consulapi.DefaultConfig()
	//config.Address = "localhost"
	fmt.Println("defautl config : ", config)
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	err = client.Agent().ServiceDeregister(Id)
	if err != nil {
		log.Fatal("register server error : ", err)
	}

}

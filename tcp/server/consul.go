package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	ID      string
	Name    string
	Address string
	Port    int
	//Check
}

//http://192.168.232.204:8550/v1/agent/service/register
type Consul struct {
	Host string
	Port int
}

var consul *Consul

func init() {
	consul = &Consul{
		Host: "http://192.168.5.200",
		Port: 8500,
	}
}

func RegisterServer(server *Server) {
	regRequest := "/v1/agent/service/register"
	data, _ := json.Marshal(server)
	r, err := http.NewRequest("PUT", consul.Host+":"+strconv.Itoa(consul.Port)+regRequest, bytes.NewReader(data))
	if err != nil {
		fmt.Println("error:", err)
	}
	res, err := http.DefaultClient.Do(r)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(res)
}

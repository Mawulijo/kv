package cmd

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

func GetConn(baseURL string, args []string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s", connUrl, args[0]), nil)
	if err != nil {
		log.Println("Could not connect to kv server")
	}
	return resp
}

func GetPing(baseUrl string) *grequests.Response {
	resp, err := grequests.Get(baseUrl, nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func SetKeyValue(baseUrl string, args []string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s/%s", baseUrl, args[0], args[1]), nil)
	if err != nil {
		log.Println("Could not create record")
	}
	return resp
}
func GetValue(baseUrl string, args string) *grequests.Response {
	resp, err := grequests.Get(fmt.Sprintf("%s/%s", baseUrl, args), nil)
	if err != nil {
		log.Println("Could not get record")
	}
	return resp
}

func DelKeyValue(baseUrl string, args string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s", baseUrl, args), nil)
	if err != nil {
		log.Println("Could not delete record")
	}
	return resp
}


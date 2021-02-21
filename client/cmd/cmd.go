package cmd

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

func GetConn(connUrl string, args []string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s", connUrl, args[0]), nil)
	if err != nil {
		log.Println("Could not connect to kv server")
	}
	return resp
}

func GetPing(pingUrl string) *grequests.Response {
	resp, err := grequests.Get(pingUrl, nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func SetKeyValue(setUrl string, args []string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s/%s", setUrl, args[0], args[1]), nil)
	if err != nil {
		log.Println("Could not create record")
	}
	return resp
}
func GetValue(getUrl string, args string) *grequests.Response {
	resp, err := grequests.Get(fmt.Sprintf("%s/%s", getUrl, args), nil)
	if err != nil {
		log.Println("Could not get record")
	}
	return resp
}

func DelKeyValue(delUrl string, args string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s", delUrl, args), nil)
	if err != nil {
		log.Println("Could not delete record")
	}
	return resp
}


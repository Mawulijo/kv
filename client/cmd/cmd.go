package cmd

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
)


func GetPing(baseUrl string) *grequests.Response {
	resp, err := grequests.Get(baseUrl, nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func SetKeyValue(baseUrl string, args []string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s/%s/set", baseUrl, args[0], args[1]), nil)
	if err != nil {
		log.Println("Could not create record")
	}
	return resp
}
func GetValue(baseUrl string, args string) *grequests.Response {
	resp, err := grequests.Get(fmt.Sprintf("%s/%s/get", baseUrl, args), nil)
	if err != nil {
		log.Println("Could not get record")
	}
	return resp
}
func GetAllKeys(baseUrl string) *grequests.Response {
	resp, err := grequests.Get(fmt.Sprintf("%s/all",baseUrl), nil)
	if err != nil {
		log.Println("Could not get all records")
	}
	return resp
}

func UpdateKeys(baseUrl string, args []string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s/%s/update", baseUrl, args[0], args[1]), nil)
	if err != nil {
		log.Println("Could not update record")
	}
	return resp
}

func DelKeyValue(baseUrl string, args string) *grequests.Response {
	resp, err := grequests.Delete(fmt.Sprintf("%s/%s/delete", baseUrl, args), nil)
	if err != nil {
		log.Println("Could not delete record")
	}
	return resp
}

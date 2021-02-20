package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/urfave/cli"
	"log"
	"os"
)

func getPing(pingUrl string) *grequests.Response {
	resp, err := grequests.Get(pingUrl, nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func setKeyValue(setUrl string, args []string) *grequests.Response {
	resp, err := grequests.Get(fmt.Sprintf("%s/%s/%s",setUrl, args[0], args[1]), nil)
	if err != nil {
		log.Println("Could not create record")
	}
	return resp
}

func getValue(getUrl string, args []string) *grequests.Response {
	resp, err := grequests.Get(fmt.Sprintf("%s/%s",getUrl, args[0]), nil)
	if err != nil {
		log.Println("Could not get record")
	}
	return resp
}

func delKeyValue(delUrl string, args []string) *grequests.Response {
	resp, err := grequests.Post(fmt.Sprintf("%s/%s",delUrl, args[0]), nil)
	if err != nil {
		log.Println("Could not delete record")
	}
	return resp
}

func main() {
	app := cli.NewApp()
	// define command for our client
	app.Commands = []cli.Command{
		{
			Name:    "ping",
			Aliases: []string{"p"},
			Usage:   "Test if server connection is alive. [Usage]: kv ping",
			Action: func(c *cli.Context) error {
				var pingUrl = "http://localhost:1024/kv/v1/store/ping"
				resp := getPing(pingUrl)
				log.Println(resp)
				return nil
			},
		},
		{
			Name:    "set",
			Aliases: []string{"s"},
			Usage:   "Sets a value for a given key. [Usage]: kv set A 42",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					args := c.Args()
					var setUrl = "http://localhost:1024/kv/v1/store/set"
					resp := setKeyValue(setUrl, args)
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Gets a value for a given key. [Usage]: kv get A",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					args := c.Args()
					var getUrl = "http://localhost:1024/kv/v1/store/get"
					resp := getValue(getUrl, args)
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Deletes a record from the store. [Usage]: kv delete A",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					args := c.Args()
					var delUrl = "http://localhost:1024/kv/v1/store/delete"
					resp := delKeyValue(delUrl, args)
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments. See -h to see help")
				}
				return nil
			},
		},
	}

	app.Version = "1.0"
	app.Run(os.Args)
}

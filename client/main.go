package main

import (
	"bufio"
	"fmt"
	"kv/cmd"
	"log"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmdString string) error {
	//app := cli.NewApp()
	cmdString = strings.TrimSuffix(cmdString, "\n")
	arrCmdString := strings.Fields(cmdString)

	switch arrCmdString[0] {
	case "--help", "-h":
		{
		fmt.Println(`
NAME:
kv - A Key-Value datastore CLI application
			
USAGE:
kv [global options] command [command options] [arguments...]
			
VERSION:
1.0
			
COMMANDS:
ping, p     Test if server connection is alive.
set, s      Sets a value for a given key. [Usage]: set A 42
get, g      Gets a value for a given key. [Usage]: get A
delete, d   Deletes a record from the store. [Usage]: delete A
help, h     Shows a list of commands or help for one command
			
GLOBAL OPTIONS:
-- connect, -c    connect to kv server
--help, -h     show help
--version, -v  print the version`)}

	//used ring because ping already exists OS
	case "ring", "r":
		{
			var pingUrl = "http://localhost:1024/kv/v1/store/ping"
			resp := cmd.GetPing(pingUrl)
			log.Println(resp)
		}

		cmd := exec.Command(arrCmdString[0], arrCmdString[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}

	// define commands for our client
	//app.Commands = []cli.Command{
	//	{
	//		Name:    "ping",
	//		Aliases: []string{"p"},
	//		Usage:   "Test if server connection is alive. [Usage]: kv ping",
	//		Action: func(c *cli.Context) error {
	//			var pingUrl = "http://localhost:1024/kv/v1/store/ping"
	//			resp := getPing(pingUrl)
	//			log.Println(resp)
	//			return nil
	//		},
	//	},
	//	{
	//		Name:    "connect",
	//		Aliases: []string{"c"},
	//		Usage:   "Connect to kv server. [Usage]: kv connect 127.0.0.1",
	//		Action: func(c *cli.Context) error {
	//			if c.NArg() > 0 {
	//				args := c.Args()
	//				var pingUrl = "http://localhost:1024/kv/v1/store"
	//				resp := getConn(pingUrl, args)
	//				log.Println(resp)
	//			} else {
	//				log.Println("Please give sufficient arguments. See -h to see help")
	//			}
	//			return nil
	//		},
	//	},
	//	{
	//		Name:    "set",
	//		Aliases: []string{"s"},
	//		Usage:   "Sets a value for a given key. [Usage]: kv set A 42",
	//		Action: func(c *cli.Context) error {
	//			if c.NArg() > 0 {
	//				args := c.Args()
	//				var setUrl = "http://localhost:1024/kv/v1/store/set"
	//				resp := setKeyValue(setUrl, args)
	//				log.Println(resp.String())
	//			} else {
	//				log.Println("Please give sufficient arguments. See -h to see help")
	//			}
	//			return nil
	//		},
	//	},
	//	{
	//		Name:    "get",
	//		Aliases: []string{"g"},
	//		Usage:   "Gets a value for a given key. [Usage]: kv get A",
	//		Action: func(c *cli.Context) error {
	//			if c.NArg() > 0 {
	//				args := c.Args()
	//				var getUrl = "http://localhost:1024/kv/v1/store/get"
	//				resp := getValue(getUrl, args)
	//				log.Println(resp.String())
	//			} else {
	//				log.Println("Please give sufficient arguments. See -h to see help")
	//			}
	//			return nil
	//		},
	//	},
	//	{
	//		Name:    "delete",
	//		Aliases: []string{"d"},
	//		Usage:   "Deletes a record from the store. [Usage]: kv delete A",
	//		Action: func(c *cli.Context) error {
	//			if c.NArg() > 0 {
	//				args := c.Args()
	//				var delUrl = "http://localhost:1024/kv/v1/store/delete"
	//				resp := delKeyValue(delUrl, args)
	//				log.Println(resp.String())
	//			} else {
	//				log.Println("Please give sufficient arguments. See -h to see help")
	//			}
	//			return nil
	//		},
	//	},
	//}
	//app.Version = "1.0"
	//app.Run(os.Args)
	return nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		_ = runCommand(cmdString)
	}

}

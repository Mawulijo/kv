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
	//used ring because ping already exists OS
	case "ring", "r":
		{
			var pingUrl = "http://localhost:1024/kv/v1/store/ping"
			resp := cmd.GetPing(pingUrl)
			log.Println(resp)

			cmd := exec.Command(arrCmdString[0], arrCmdString[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			return cmd.Run()
		}
	case "set", "s":
		{
			args := []string{arrCmdString[1], arrCmdString[2]}
			var setUrl = "http://localhost:1024/kv/v1/store/set"
			resp := cmd.SetKeyValue(setUrl, args)
			log.Println(resp.String())
		}
	case "get", "g":
		{
			var getUrl = "http://localhost:1024/kv/v1/store/get"
			resp := cmd.GetValue(getUrl, arrCmdString[1])
			log.Println(resp.String())
		}
		case "delete", "d":
		{
			var detUrl = "http://localhost:1024/kv/v1/store/delete"
			resp := cmd.DelKeyValue(detUrl, arrCmdString[1])
			log.Println(resp.String())
		}
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

	}
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

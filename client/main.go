package main

import (
	"bufio"
	"fmt"
	"io"
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
	case "help", "h":
		{
			fd, _ := os.Open("helpFile")
			if _, err := io.Copy(os.Stdout, fd); err != nil {
				log.Fatal(err)
			}
		}
		case "exit", "q":
		{
			os.Exit(1)
		}

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

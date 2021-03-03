package main

import (
	"bufio"
	"fmt"
	"io"
	"kv/cmd"
	"log"
	"os"
	"strings"
)

func runCommand(cmdString string) error {
	baseUrl := "http://localhost:1024/v1/kv"
	cmdString = strings.TrimSuffix(cmdString, "\n")
	arrCmdString := strings.Fields(cmdString)

	switch arrCmdString[0] {
	//used ring because ping already exists in OS
	case "ring", "r":
		{
			resp := cmd.GetPing(baseUrl)
			log.Println(resp)
		}
	case "set", "s":
		{
			args := []string{arrCmdString[1], arrCmdString[2]}
			resp := cmd.SetKeyValue(baseUrl, args)
			log.Println(resp.String())
		}
	case "get", "g":
		{
			resp := cmd.GetValue(baseUrl, arrCmdString[1])
			log.Println(resp.String())
		}
	case "get all", "ga":
		{
			resp := cmd.GetAllKeys(baseUrl)
			log.Println(resp.String())
		}
	case "update", "u":
		{
			args := []string{arrCmdString[1], arrCmdString[2]}
			resp := cmd.UpdateKeys(baseUrl, args)
			log.Println(resp.String())
		}
	case "delete", "d":
		{
			resp := cmd.DelKeyValue(baseUrl, arrCmdString[1])
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

	default: {
		log.Println("Invalid Command")
	}
	}
	return nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("KV > ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		_ = runCommand(cmdString)
	}

}

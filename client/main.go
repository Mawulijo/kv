package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"kv/cmd"
	"log"
	"os"
	"strings"
)

func runCommand(cmdString, addr string) error {
	baseUrl := fmt.Sprintf("http://%s:1024/v1/kv", addr)
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

	default:
		{
			log.Println("Unknown Command")
		}
	}
	return nil
}

func main() {
	serverAddr := "127.0.0.1"
	h := flag.String("h", "", "Provide this flag to connect to kv server")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	if string(*h) != serverAddr {
		fmt.Println("Invalid Server Address")
		return
	}

	for {
		fmt.Print("KV > ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString, string(*h))
		if err != nil {
			return
		}
	}
}

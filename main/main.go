package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	commandArgs := os.Args

	firstArgs := commandArgs[1]

	switch firstArgs {
	case "create":
		if len(commandArgs) <= 2 {
			fmt.Print("No directory name, create failed.")
			return
		}
	case "startserver":
		flag.Parse()
	}
}

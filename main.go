package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/chaos-io/go-verse/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the go-verse programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command!\n")
	repl.Start(os.Stdin, os.Stdout)
}

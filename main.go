package main

import (
	"fmt"
	"intergo/pkg/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Welcome", user.Name)
	fmt.Println("Type some commands!")
	repl.Begin(os.Stdin, os.Stdout)
}

package main

import (
	"barracuda/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Barracuda programming language!\n", user.Username)
	fmt.Println("⬇️ Start")

	repl.Start(os.Stdin, os.Stdout)
}

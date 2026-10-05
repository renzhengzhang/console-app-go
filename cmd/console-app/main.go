package main

import (
	"fmt"
	"os"

	"renzheng.me/console-app-go/internal/app"
)

func main() {
	if err := app.Run(os.Stdin, os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

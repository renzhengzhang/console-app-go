package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Run(in io.Reader, out io.Writer, args []string) error {
	var name string
	if len(args) > 0 {
		name = args[0]
	} else {
		fmt.Fprint(out, "Please enter your name: ")

		scanner := bufio.NewScanner(in)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return err
			}

			return fmt.Errorf("Must enter your name")
		}
		name = scanner.Text()
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("Name cannot be empty")
	}

	_, err := fmt.Fprintf(out, "Hello, %s!\n", name)
	return err
}

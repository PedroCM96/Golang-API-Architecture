package commands

import "fmt"

type DummyCommand struct{}

func (c DummyCommand) Exec(args []string) {
	arg := args[0]
	fmt.Println(arg)
}

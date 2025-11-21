package main

import (
	"canvas/cmd"
	"fmt"
)

func main() {
	err := cmd.NewListener()
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"canvas/cmd"
	"fmt"
)

func main() {
	err := cmd.NewListener()
	fmt.Println(err)
}

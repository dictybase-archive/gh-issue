package commands

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

func Hello() {
	fmt.Println("hello")
}
func TestFunction(c *cli.Context) {
	fmt.Println("testing out the server")
}

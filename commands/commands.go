package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

func TestFunction(c *cli.Context) error {
	fmt.Println("testing out the server")
	return nil
}

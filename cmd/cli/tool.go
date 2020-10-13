package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("Oh, hello")
	
	(&cli.App{}).Run(os.Args)
}

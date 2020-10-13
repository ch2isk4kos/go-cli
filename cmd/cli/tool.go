package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "test",
		Usage: "a test command",
		Action: func(c *cli.Context) error {
			fmt.Println("Testing...Testing...")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	} 
}

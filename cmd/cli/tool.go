package main

import (
	// "fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
  // (&cli.App{}).Run(os.Args)
	
	// app := &cli.App{
	// 	Name: "application name",
	// 	Usage: "application description",
	// 	Action: func(c *cli.Context) error {
	// 		fmt.Println("write something here to help navigate")
	// 		fmt.Printf("%q\n", c.Args().Get(1))
	// 		return nil
	// 	},
	// }

	app := cli.NewApp()
	app.Name = "name of application"
	app.Usage = "application description"
	// app.Action = func(c *cli.Context) error {
	// 			fmt.Println("write something here to help navigate")
	// 			return nil
	// }

	app.Flags = []cli.Flag {		
		&cli.StringFlag{
			Name: "flag-name",
			Value: "default-value",
		},
	}
	// fmt.Println("flags: ")
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

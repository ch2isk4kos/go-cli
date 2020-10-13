package main

import (
	"fmt"
	"log"
	"net"
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

	// instantiate application
	app := cli.NewApp()
	app.Name = "name of application"
	app.Usage = "application description"

	// flags
	// fs := []cli.Flag {		
	// 	&cli.StringFlag{
	// 		Name: "flag-name",
	// 		Value: "default-value",
	// 	},
	// }

	// fmt.Println("flags: ", fs)

	app.Commands = []*cli.Command{
		{
			Name: "ghost",
			Usage: "get network host of url provided",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "host",
					Value: "geeksforgeeks.org",
				},
			},
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					log.Fatal(err)
				}

				// fmt.Println("net server: ", ns)

				for i := 0; i < len(ns); i++ {
					fmt.Println("network host: ", ns[i].Host)	
				}
				return nil
			},
		},
		{
			Name: "gip",
			Usage: "get ip of url provided",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "flag-name",
					Value: "google.com",
				},
			},
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS("com")
				if err != nil {
					log.Fatal(err)
				}

				// fmt.Println("net server: ", ns)

				for i := 0; i < len(ns); i++ {
					fmt.Println("network host: ", ns[i].Host)	
				}
				return nil
			},
		},
	}

	// Run Application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}



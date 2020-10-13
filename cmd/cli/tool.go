package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

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
	app.Name = "Go CLI Lookup"
	app.Usage = "Command Line Interface tool that handles automated network queries."

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
			Name: "gname",
			Usage: "get NS records of domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "host",
					Value: "https://www.golang.org",
				},
			},
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS("host")
				if err != nil {
					log.Fatal(err)
				}

				// for i := 0; i < len(ns); i++ {
				// 	fmt.Println("Host: ", ns[i].Host)	
				// }

				for _, v := range ns {
					fmt.Println("Host: ", v.Host)	
				}
				return nil
			},
		},
		{
			Name: "ghost",
			Usage: "get network host of domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "net",
					Value: "https://www.golang.org",
				},
			},
			Action: func(c *cli.Context) error {
				h, err := net.LookupHost(c.String("net"))
				// h, err := net.LookupHost("com")
				if err != nil {
					log.Fatal(err)
				}

				// fmt.Println("net server: ", ns)

				for i := 0; i < len(h); i++ {
					hosts := strings.SplitAfter(h[i], " ") 
					fmt.Println("network host: ", hosts)	
				}
				return nil
			},
		},
		{
			Name: "gip",
			Usage: "get ip of domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "host",
				},
			},
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					log.Fatal(err)
				}

				// fmt.Println("net server: ", ip)

				for i := 0; i < len(ip); i++ {
					fmt.Println("host ip: ", ip[i])	
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



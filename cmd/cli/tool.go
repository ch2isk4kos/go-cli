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
	// var input string
	// fmt.Scanf("%s", input)

	app.Commands = []*cli.Command{
		{
			Name: "gname",
			Usage: "get ns records of domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "host",
					Value: "NS records",
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

				for i, v := range ns {
					fmt.Printf("Host %v: %v\n", i+1, v.Host)	
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
					Value: "network host",
				},
			},
			Action: func(c *cli.Context) error {
				h, err := net.LookupHost(c.String("net"))
				if err != nil {
					log.Fatal(err)
				}

				for i := 0; i < len(h); i++ {
					hosts := strings.SplitAfter(h[i], " ")
					if len(hosts) == 0 {
						fmt.Println("No Hosts Found.")
					} 
					fmt.Println("Network Host: ", hosts)	
				}
				return nil
			},
		},
		{
			Name: "gaddr",
			Usage: "get ip address of domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "ipa",
					Value: "IP address",
				},
			},
			Action: func(c *cli.Context) error {
				ips, err := net.LookupIP(c.String("ipa"))

				if err != nil {
					log.Fatal(err)
				}

				// for i := 0; i < len(ip); i++ {
				// 	fmt.Println("host ip: ", ip[i])	
				// }

				for i := 0; i < len(ips); i++ {
					addr := strings.SplitAfter(ips[i].String(), " ")
					if len(addr) == 0 {
						fmt.Println("No Hosts Found.")
					} 
					fmt.Println("IP Address: ", addr)	
				}
				return nil
			},
		},
		{
			Name: "gport",
			Usage: "get tcp port of domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "num",
					Value: "port",
				},
			},
			Action: func(c *cli.Context) error {
				pTCP, err := net.LookupPort("tcp", "domain")
				pUDP, err := net.LookupPort("udp", "domain")
				
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("TCP Port: %d\n", pTCP)
				fmt.Printf("UDP Port: %d\n", pUDP)

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



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
	// instantiate application
	app := cli.NewApp()
	app.Name = "Go CLI Lookup"
	app.Usage = "Command Line Interface tool that handles automated network queries."

	app.Commands = []*cli.Command{
		{
			Name: "gname",
			Usage: "Get NameServer Records of Domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "host",
					Value: "google.com",
				},
			},
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					log.Fatal(err)
				}

				for i, v := range ns {
					fmt.Printf("Host %v: %v\n", i+1, v.Host)	
				}
				return nil
			},
		},
		{
			Name: "ghost",
			Usage: "Get Network Host of Domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "net",
					Value: "www.google.com",
				},
			},
			Action: func(c *cli.Context) error {
				h, err := net.LookupHost(c.String("net"))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("")

				for i := 0; i < len(h); i++ {
					hosts := strings.SplitAfter(h[i], " ")
					if len(hosts) == 0 {
						fmt.Println("No Hosts Found.")
					}
					fmt.Println("Network Host: ", hosts)	
				}
				fmt.Println("")
				return nil
			},
		},
		{
			Name: "gaddr",
			Usage: "Get IP Address of Domain",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "ip",
					Value: "google.com",
				},
			},
			Action: func(c *cli.Context) error {
				ips, err := net.LookupIP(c.String("ip"))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("")
				
				for i := 0; i < len(ips); i++ {
					addr := strings.SplitAfter(ips[i].String(), " ")
					if len(addr) == 0 {
						fmt.Println("No Hosts Found.")
					} 
					fmt.Println("IP Address: ", addr)	
				}
				fmt.Println("")
				return nil
			},
		},
		{
			Name: "grdda",
			Usage: "Get Reverse Lookup for Address",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "rev",
					Value: "8.8.8.8",
				},
			},
			Action: func(c *cli.Context) error {
				addrs, err := net.LookupAddr(c.String("rev"))

				if err != nil {
					log.Fatal(err)
				}

				for _, v := range addrs {
					fmt.Printf("\nReverse Address: %s\n", v)
					fmt.Println("")
				}
				return nil
			},
		},
		{
			Name: "gport",
			Usage: "Get Port Number of Utility",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "ntwk",
					Value: "tcp",
				},
				&cli.StringFlag{
					Name: "srvc",
					Value: "domain",
				},
			},
			Action: func(c *cli.Context) error {
				port, err := net.LookupPort(c.String("ntwk"), c.String("srvc"))
				if err != nil {
					log.Fatal(err)
				}
				
				fmt.Printf("Port: %d\n\n", port)
				return nil
			},
		},
		{
			Name: "gmx",
			Usage: "get mx records of domain",
			Flags: []cli.Flag {
				&cli.StringFlag{
					Name: "host",
					Value: "google.com",
				},
			},
			Action: func(c *cli.Context) error {
				mxr, err := net.LookupMX(c.String("host"))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("")

				for _, mx := range mxr {
					fmt.Printf("Host: %s\tPref: %v\n", mx.Host, mx.Pref)
				}

				fmt.Println("")
				return nil
			},
		},
		{
			Name: "gsrv",
			Usage: "get srv query of service - protocol - domain",
			Flags: []cli.Flag {
				&cli.StringFlag{
					Name: "cn",
					Value: "cname",
				},
				&cli.StringFlag{
					Name: "proto",
					Value: "protocol",
				},
				&cli.StringFlag{
					Name: "host",
					Value: "domain",
				},
			},
			Action: func(c *cli.Context) error {
				cname, srvs, err := net.LookupSRV(c.String("cn"), c.String("proto"), c.String("host"))
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("\nCNAME: %s\n\n", cname)

				for _, srv := range srvs {
					fmt.Printf("TARGET: %v\tPORT: %v\tPRIORITY: %v\tWEIGHT: %v\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
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



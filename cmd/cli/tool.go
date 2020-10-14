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
			Usage: "get network host of domain",
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
			Name: "grdda",
			Usage: "get reverse lookup for address",
			Flags: []cli.Flag {		
				&cli.StringFlag{
					Name: "rev",
					Value: "reverse address",
				},
			},
			Action: func(c *cli.Context) error {
				addrs, err := net.LookupAddr(c.String("rev"))

				if err != nil {
					log.Fatal(err)
				}

				for _, v := range addrs {
					fmt.Printf("Address: %s\n", v)	
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
		{
			Name: "gmx",
			Usage: "get mx records of domain",
			Flags: []cli.Flag {
				&cli.StringFlag{
					Name: "host",
					Value: "mx records",
				},
			},
			Action: func(c *cli.Context) error {
				mxr, err := net.LookupMX(c.String("host"))
				if err != nil {
					log.Fatal(err)
				}

				for _, mx := range mxr {
					fmt.Printf("Host: %s\nPref: %v", mx.Host, mx.Pref)
				}
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



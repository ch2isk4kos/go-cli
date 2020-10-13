package main

import ( 
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli" 
)

func main() {
	// ENTRY POINT
	app := cli.NewApp()
	app.Name = "Web Query CLI"
	app.Usage = "query IPs, CNAMEs, MX records and Name Servers"	
	
	// CREATE FLAGS
	fs := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Usage: "google lookup",
			Value: "https://www.google.com",
		},
	}

	// CREATE COMMANDS
	app.Commands = []cli.Command{
		{
			Name: "ns",
			Usage: "name server lookup for specified host",
			Flags: fs,

			// THE CODE THAT WILL EXECUTE
			Action: func(c *cli.Context) error{
				ns, err := net.LookupNS(c.String("url"))
				if err != nil {
					fmt.Println("err: ", err)
					return err
				}
				
				// LOG TO CONSOLE
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},	
	}

	// START APPLICATION
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

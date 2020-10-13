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
	app.Name = "Web Query CLII"
	app.Usage = "query  IPs, CNAMEs, MX records and Name Servers!"
	
	
	// CREATE FLAGS
	fs := []cli.Flag{
		cli.StringFlag{
			Name: "host"
			Value: "Chris Kakos"
		}
	}

	// CREATE COMMANDS
	app.Commands = [].cli.Commands{
		
	}
	
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

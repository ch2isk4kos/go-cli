package main

import ( 
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli" 
)

func main() {
	app := cli.NewApp()
	app.Name = "Web Query CLII"
	app.Usage = "query  IPs, CNAMEs, MX records and Name Servers!"
	
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

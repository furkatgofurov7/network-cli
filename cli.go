package main

import (
	"log"
	"net"
	"os"
	"fmt"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "It let's you query IP addresses, CNAMEs, MX records, Name Servers and Ports!"

	myFlags := []cli.Flag{
		&cli.StringFlag {
			Name: "host",
			Value: "github.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a Particular Host",
			Flags: myFlags,
			Action: func (c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name: "ip",
			Usage: "Looks up the IP addresses for a Particular Host",
			Flags: myFlags,
			Action: func (c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}
				for i :=0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name: "cname",
			Usage: "Looks up the CNAME for a Particular Host",
			Flags: myFlags,
			Action: func (c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name: "mx",
			Usage: "Looks up the MX records for a Particular Host",
			Flags: myFlags,
			Action: func (c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i :=0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
		{
			Name: "port",
			Usage: "Looks up the port for the given network and service",
			Flags: myFlags,
			Action: func (c *cli.Context) error {
				// LookupPort receives 2 args as follows:
				// LookupPort(network, service string)
				// network can be: "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6"
				// service arg can be any protocol supported by above provided network options.
				port, err := net.LookupPort("tcp", "https")
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println("https port:", port)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}




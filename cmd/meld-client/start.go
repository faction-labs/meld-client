package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/factionlabs/meld-agent/pkg/rust"
)

var rustStartCmd = cli.Command{
	Name:   "start-rust",
	Usage:  "start rust server",
	Action: rustStartCmdAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "hostname",
			Usage: "rust server hostname",
			Value: "Meld Server",
		},
		cli.StringFlag{
			Name:  "description",
			Usage: "rust server description",
			Value: "Default Meld Rust Server",
		},
		cli.StringFlag{
			Name:  "identity",
			Usage: "set rust identity",
			Value: "meld_server",
		},
		cli.StringFlag{
			Name:  "server-ip",
			Usage: "server ip address",
			Value: "0.0.0.0",
		},
		cli.IntFlag{
			Name:  "server-port",
			Usage: "port of server",
			Value: 28015,
		},
		cli.StringFlag{
			Name:  "rcon-ip",
			Usage: "rcon ip address",
			Value: "0.0.0.0",
		},
		cli.IntFlag{
			Name:  "rcon-port",
			Usage: "rcon rcon",
			Value: 28016,
		},
		cli.StringFlag{
			Name:  "rcon-password",
			Usage: "rcon password",
			Value: "meld",
		},
		cli.IntFlag{
			Name:  "maxplayers",
			Usage: "max players",
			Value: 16,
		},
		cli.IntFlag{
			Name:  "seed",
			Usage: "level seed",
			Value: 5050,
		},
		cli.IntFlag{
			Name:  "worldsize",
			Usage: "size of the world",
			Value: 4096,
		},
		cli.IntFlag{
			Name:  "save-interval",
			Usage: "time between saves",
			Value: 300,
		},
		cli.StringFlag{
			Name:  "url",
			Usage: "set info url",
			Value: "https://github.com/factionlabs/meld",
		},
	},
}

func rustStartCmdAction(c *cli.Context) {
	addr := c.GlobalString("addr")

	client, err := getClient(addr)
	if err != nil {
		log.Fatal(err)
	}

	args := &rust.StartArgs{
		Hostname:     c.String("hostname"),
		Description:  c.String("description"),
		Identity:     c.String("identity"),
		ServerIP:     c.String("server-ip"),
		ServerPort:   c.Int("server-port"),
		RconIP:       c.String("rcon-ip"),
		RconPort:     c.Int("rcon-port"),
		RconPassword: c.String("rcon-password"),
		MaxPlayers:   c.Int("maxplayers"),
		Level:        "Procedural Map",
		Seed:         c.Int("seed"),
		WorldSize:    c.Int("worldsize"),
		SaveInterval: c.Int("save-interval"),
		URL:          c.String("url"),
	}

	var pid int

	if err := client.Call("Meld.StartRust", args, &pid); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rust Server Start: pid=%d\n", pid)
}

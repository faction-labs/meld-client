package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/factionlabs/meld-agent/pkg/rust"
)

var rustStartCmd = cli.Command{
	Name:   "rust-start",
	Usage:  "start rust server",
	Action: rustStartCmdAction,
}

func rustStartCmdAction(c *cli.Context) {
	addr := c.GlobalString("addr")

	client, err := getClient(addr)
	if err != nil {
		log.Fatal(err)
	}

	args := &rust.StartArgs{
		Hostname:     "meld server",
		Description:  "meld rust server",
		Identity:     "meld_server",
		ServerIP:     "0.0.0.0",
		ServerPort:   28015,
		RconIP:       "0.0.0.0",
		RconPort:     28016,
		RconPassword: "meld",
		MaxPlayers:   16,
		Level:        "Procedural Map",
		Seed:         5050,
		WorldSize:    4096,
		SaveInterval: 300,
		URL:          "https://github.com/factionlabs/meld",
	}

	var success bool

	if err := client.Call("Meld.StartRust", args, &success); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rust Server Start: success=%v", success)
}

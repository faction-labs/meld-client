package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var rustStopCmd = cli.Command{
	Name:   "stop-rust",
	Usage:  "stop rust server",
	Action: rustStopCmdAction,
}

func rustStopCmdAction(c *cli.Context) {
	addr := c.GlobalString("addr")

	client, err := getClient(addr)
	if err != nil {
		log.Fatal(err)
	}

	var success bool

	if err := client.Call("Meld.StopRust", 0, &success); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rust Server Stop: successful=%v\n", success)
}

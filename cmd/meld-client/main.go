package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "meld-client"
	app.Usage = "meld agent client"
	app.Author = "faction labs"
	app.Version = "0.1.0"
	app.Email = ""
	app.Commands = []cli.Command{
		installSteamCmd,
		installRustCmd,
		installOxideCmd,
		versionCmd,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr, a",
			Usage: "meld agent address",
			Value: "127.0.0.1:9000",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

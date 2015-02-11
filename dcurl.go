package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dcurl"
	app.Version = Version
	app.Usage = "Depleted curl like command line tool"
	app.Author = "Daisuke Miyamoto"
	app.Email = "dai.0304@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}

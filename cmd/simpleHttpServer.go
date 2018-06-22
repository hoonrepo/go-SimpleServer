package cmd

import (
	"fmt"
	"github.com/hoonrepo/go-SimpleServer/params"
	"github.com/hoonrepo/go-SimpleServer/utils"
	"gopkg.in/urfave/cli.v2"
	"os"
	"os/signal"
)

var (
	app      = &cli.App{}
	commands = []*cli.Command{
		utils.HttpCmd,
	}
)

func init() {
	app.Version = params.Version
	app.Commands = commands

}

func main() {

	if app == nil {
		os.Exit(1)
	}

	sn := make(chan os.Signal, 1)
	signal.Notify(sn, os.Interrupt, os.Kill)

	go func() {
		for sig := range sn {
			fmt.Printf("%v simplehttp has been stopped!\n", sig)
			os.Exit(0)
		}
	}()

	if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

}

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	app   = NewApp()
	sever = new(srv)
)

func init() {
	app.Action = cli.ShowAppHelp
	app.Commands = []cli.Command{
		startCommand,
	}

	app.Flags = []cli.Flag{
		HTTPHostFlag,
		HTTPPortFlag,
		MongoHostFlag,
		MongoPortFlag,
		MongoDatabaseFlag,
		TokenKeyFlag,
		CookieNameFlag,
	}

}

// Start ...
func Start(ctx *cli.Context) error {
	if err := sever.loadConfig(ctx); err != nil {
		return err
	}

	if err := sever.connectMongo(); err != nil {
		return err
	}

	if err := sever.loadLogger(); err != nil {
		return err
	}

	if err := sever.loadRepository(); err != nil {
		return err
	}

	if err := sever.loadAuthenticator(); err != nil {
		return err
	}

	if err := sever.loadDomain(); err != nil {
		return err
	}

	if err := sever.loadBucket(); err != nil {
		return err
	}

	if err := sever.loadHTTPServer(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

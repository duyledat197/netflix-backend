package main

import (
	"github.com/duyledat197/netfix-backend/utils"
	"github.com/urfave/cli"
)

var (

	// HTTPPortFlag ...
	HTTPHostFlag = cli.StringFlag{
		Name:   "http.host",
		Usage:  "Host listen",
		EnvVar: "HTTP_HOST",
		Value:  "localhost",
	}

	// HTTPPortFlag ...
	HTTPPortFlag = cli.StringFlag{
		Name:   "http.port",
		Usage:  "Port listen",
		EnvVar: "HTTP_PORT",
		Value:  "10010",
	}
	// MongoHostFlag ...
	MongoHostFlag = cli.StringFlag{
		Name:   "mongo.host",
		Usage:  "Mongo host",
		EnvVar: "MONGO_HOST",
		Value:  "localhost",
	}
	// MongoPortFlag ...
	MongoPortFlag = cli.StringFlag{
		Name:   "mongo.port",
		Usage:  "Mongo port",
		EnvVar: "MONGO_PORT",
		Value:  "27017",
	}
	// MongoDatabaseFlag ...
	MongoDatabaseFlag = cli.StringFlag{
		Name:   "mongo.database",
		Usage:  "Mongo database",
		EnvVar: "MONGO_DATABASE",
		Value:  "qanda",
	}

	// TokenKeyFlag ...
	TokenKeyFlag = cli.StringFlag{
		Name:   "token.key",
		Usage:  "token key",
		EnvVar: "TOKEN_KEY",
		Value:  utils.RandStringBytes(32),
	}
	// CookieNameFlag ...
	CookieNameFlag = cli.StringFlag{
		Name:   "cookie.name",
		Usage:  "cookie name",
		EnvVar: "COOKIE_NAME",
		Value:  "c5token",
	}
)

var (
	startCommand = cli.Command{
		Action:      migrateFlags(Start),
		Name:        "start",
		Usage:       "Bootstrap and start worker server",
		ArgsUsage:   "<genesisPath>",
		Flags:       []cli.Flag{},
		Category:    "Crawler Worker",
		Description: `Used to start crawler worker, clone data from omada cloud`,
	}
)

// NewApp creates an app with sane defaults.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Action = cli.ShowAppHelp
	app.Name = "QandA"
	app.Author = "Le Duy Dat"
	app.Email = "duyledat197@gmail.com"
	app.Usage = "question and answer"
	return app
}

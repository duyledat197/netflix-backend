package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// Config ...
type Config struct {
	HTTP ConnAddress
	GRPC ConnAddress

	Mongo Mongo

	// for authentication
	CookieName      string
	TokenKey        string
	TokenExpiration int
}

// Mongo ...
type Mongo struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (m *Mongo) getConnectString() string {
	conn := "mongodb://"
	if m.Username != "" {
		conn += fmt.Sprintf("%s:%s@", m.Username, m.Password)
	}
	conn += fmt.Sprintf("%s:%s/%s", m.Host, m.Port, m.Database)
	return conn
}

// ConnAddress ...
type ConnAddress struct {
	Host string
	Port string
}

func (c *ConnAddress) String() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func migrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			ctx.GlobalSet(name, ctx.String(name))
		}
		return action(ctx)
	}
}

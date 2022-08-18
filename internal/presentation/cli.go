package presentation

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/application"
	"github.com/MAZEN-Kenjrawi/pwd/internal/infrastructure"
)

type Cli struct {
	cmdBus infrastructure.CmdBus
}

func (cli *Cli) Signup(username, secret string) error {
	cmd := application.SignUpCommand{
		Username: username,
		Secret:   secret,
	}
	err := cli.cmdBus.Handle(cmd)
	if err != nil {
		return err
	}
	cmd2 := application.AddLoginCommand{
		ProfileUsername: username,
		Username: "Mazen7",
		Domain: "stackoverflow.com",
		Password: "123",
	}
	return cli.cmdBus.Handle(cmd2)
}

func (cli *Cli) AddLogin(profileUsername, username, domain, password string) error {
	cmd := application.AddLoginCommand{
		ProfileUsername: profileUsername,
		Username:        username,
		Domain:          domain,
		Password:        password,
	}
	err := cli.cmdBus.Handle(cmd)
	if err != nil {
		return err
	}
	return nil
}

func NewCli(bus infrastructure.CmdBus) *Cli {
	return &Cli{bus}
}

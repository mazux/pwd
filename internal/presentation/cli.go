package presentation

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/application"
	"github.com/MAZEN-Kenjrawi/pwd/internal/infrastructure"
)

type Cli struct {
	CmdBus infrastructure.CmdBus
}

func (cli *Cli) Signup(username, secret string) error {
	cmd := application.SignUpCommand{
		Username: username,
		Secret:   secret,
	}
	err := cli.CmdBus.Handle(cmd)
	if err != nil {
		return err
	}
	return nil
}

func NewCli(bus infrastructure.CmdBus) *Cli {
	return &Cli{bus}
}

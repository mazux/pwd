package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/MAZEN-Kenjrawi/pwd/internal/infrastructure"
	"github.com/MAZEN-Kenjrawi/pwd/internal/presentation"
	"golang.org/x/term"
)

func main() {
	cfg, err := infrastructure.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	c, err := infrastructure.NewContainer(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	err = c.Invoke(func(bus infrastructure.CmdBus) error {
		cli := presentation.NewCli(bus)
		username := "MAZux" // stringPrompt("Profile username: ")
		secret := "123"     // passwordPrompt("Your passowrd: ")
		return cli.Signup(username, secret)
	})

	if err != nil {
		log.Fatalln(err)
	}
}

func stringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}

	return strings.TrimSpace(s)
}

func passwordPrompt(label string) string {
	var s string
	for {
		fmt.Fprint(os.Stderr, label+" ")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(b)
		if s != "" {
			break
		}
	}
	fmt.Println()

	return s
}

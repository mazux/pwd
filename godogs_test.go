package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MAZEN-Kenjrawi/pwd/internal/infrastructure"
	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"github.com/MAZEN-Kenjrawi/pwd/internal/presentation"
	"github.com/cucumber/godog"
)

var (
	c   *infrastructure.Container
	cfg infrastructure.Config
	err error
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	cfg.Env = "test"
	cfg.Storage.Mode = infrastructure.FILE_SYSTEM_STORAGE_MODE
	cfg.Storage.Url = filepath.Join(pwd, "storage", cfg.Env)

	c, err = infrastructure.NewContainer(cfg)
	if err != nil {
		log.Fatalln(err)
	}
}

func iHaveAProfileForUsername(username string) error {
	path := filepath.Join(cfg.Storage.Url, fmt.Sprintf("%s.json", username))
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write([]byte(`{}`))

	return err
}

func iSignupUsingUsername(username string) error {
	err = c.Invoke(func(bus infrastructure.CmdBus) error {
		cli := presentation.NewCli(bus)
		return cli.Signup(username, "foo-bar")
	})

	if errors.Is(err, model.ErrorProfileAlreadyExists) {
		return nil
	}

	return err
}

func iWillGetThisError(expectedError string) error {
	if err.Error() == expectedError {
		return nil
	}

	return fmt.Errorf("unexpected error return, expected: `%s`, actual: `%s`", expectedError, err.Error())
}

func iHaveNoProfilesStored() error {
	return clearStorage()
}

func iWillHaveAStoredProfileWithUsernameAndEmptyListOfLogins(username string) error {
	path := filepath.Join(cfg.Storage.Url, fmt.Sprintf("%s.json", username))
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	p := struct {
		Username string        `json:"username"`
		Logins   []interface{} `json:"logins"`
	}{}
	err = jsonParser.Decode(&p)
	if err != nil {
		return err
	}

	if len(p.Logins) == 0 {
		return nil
	}

	return fmt.Errorf("logins list for newly signed-up user is not empty")
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		clearStorage()
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		clearStorage()
		return ctx, err
	})

	ctx.Step(`^I have a profile for username "([^"]*)"$`, iHaveAProfileForUsername)
	ctx.Step(`^I signup using username "([^"]*)"$`, iSignupUsingUsername)
	ctx.Step(`^I will get this error "([^"]*)"$`, iWillGetThisError)

	ctx.Step(`^I have no profiles stored$`, iHaveNoProfilesStored)
	ctx.Step(`^I will have a stored profile with username "([^"]*)" and empty list of logins$`, iWillHaveAStoredProfileWithUsernameAndEmptyListOfLogins)
}

func clearStorage() error {
	files, err := filepath.Glob(filepath.Join(cfg.Storage.Url, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.RemoveAll(file)
		if err != nil {
			return err
		}
	}

	return nil
}

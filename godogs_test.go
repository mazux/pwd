package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/MAZEN-Kenjrawi/pwd/internal/infrastructure"
	"github.com/MAZEN-Kenjrawi/pwd/internal/tests/testcontext"
	"github.com/cucumber/godog"
)

var (
	profileContext *testcontext.Profile
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	var cfg infrastructure.Config
	cfg.Env = "test"
	cfg.Storage.Mode = infrastructure.FILE_SYSTEM_STORAGE_MODE
	cfg.Storage.Url = filepath.Join(pwd, "storage", cfg.Env)

	c, err := infrastructure.NewContainer(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	c.Invoke(func(bus infrastructure.CmdBus) {
		profileContext = testcontext.NewProfile(bus, cfg.Storage.Url)
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		profileContext.ClearStorage()
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		profileContext.ClearStorage()
		return ctx, err
	})

	ctx.Step(`^I have a profile for username "([^"]*)"$`, profileContext.IHaveAProfileForUsername)
	ctx.Step(`^I signup using username "([^"]*)"$`, profileContext.ISignupUsingUsername)
	ctx.Step(`^I will get this error "([^"]*)"$`, profileContext.IWillGetThisError)

	ctx.Step(`^I have no profiles stored$`, profileContext.IHaveNoProfilesStored)
	ctx.Step(`^I will have a stored profile with username "([^"]*)" and empty list of logins$`, profileContext.IWillHaveAStoredProfileWithUsernameAndEmptyListOfLogins)
}

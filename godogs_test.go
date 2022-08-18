package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/MAZEN-Kenjrawi/pwd/internal/infrastructure"
	"github.com/MAZEN-Kenjrawi/pwd/internal/presentation"
	"github.com/MAZEN-Kenjrawi/pwd/internal/tests/feature"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var (
	profileFeature *feature.ProfileFeature
	opts           = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "emoji",
	}
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
		cli := presentation.NewCli(bus)
		profileFeature = feature.NewProfile(cli, cfg.Storage.Url)
	})
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	profileFeature.InitializeScenario(ctx)
}

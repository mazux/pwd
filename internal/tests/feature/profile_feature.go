package feature

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"github.com/MAZEN-Kenjrawi/pwd/internal/presentation"
	"github.com/cucumber/godog"
)

type ProfileFeature struct {
	cli                    *presentation.Cli
	storageUrl             string
	currentProfileUsername string
	lastErr                error
}

func (p *ProfileFeature) iHaveAProfileForUsername(username string) error {
	path := filepath.Join(p.storageUrl, fmt.Sprintf("%s.json", username))
	content := []byte(fmt.Sprintf(`{"username":"%s","secret":"PJkJr+wlNU1VHa4hWQuybjjVPyFzuNPcPu5MBH56scHri4UQPjvnumE7MbtcnDYhTcnxSkL9ei/bhIVrylxEwg==","logins":[]}`, username))
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	p.currentProfileUsername = username

	return err
}

func (p *ProfileFeature) iSignupUsingUsername(username string) error {
	p.lastErr = p.cli.Signup(username, "123")
	if errors.Is(p.lastErr, model.ErrorProfileAlreadyExists) {
		return nil
	}

	return p.lastErr
}

func (p *ProfileFeature) iWillGetThisError(expectedError string) error {
	if p.lastErr.Error() == expectedError {
		return nil
	}

	return fmt.Errorf("unexpected error return, expected: `%s`, actual: `%s`", expectedError, p.lastErr)
}

func (p *ProfileFeature) iHaveEmptyStorage() error {
	files, err := filepath.Glob(filepath.Join(p.storageUrl, "*"))
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

func (p *ProfileFeature) iWillHaveAStoredProfileWithUsernameAndEmptyListOfLogins(username string) error {
	path := filepath.Join(p.storageUrl, fmt.Sprintf("%s.json", username))
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	profile := struct {
		Username string        `json:"username"`
		Logins   []interface{} `json:"logins"`
	}{}
	err = jsonParser.Decode(&p)
	if err != nil {
		return err
	}

	if len(profile.Logins) == 0 {
		return nil
	}

	return fmt.Errorf("logins list for newly signed-up user is not empty")
}

func (p *ProfileFeature) itHasALoginForDomainAndUsername(domain, username string) error {
	path := filepath.Join(p.storageUrl, fmt.Sprintf("%s.json", p.currentProfileUsername))
	content := fmt.Sprintf(`{"username":"%s","secret":"PJkJr+wlNU1VHa4hWQuybjjVPyFzuNPcPu5MBH56scHri4UQPjvnumE7MbtcnDYhTcnxSkL9ei/bhIVrylxEwg==","logins":[{"username":"%s","domain":"%s", "password":"efe4mHj2qgdXG1EGYTuyFjAMx1qs0olbyUq3v2K+HA=="}]}`, p.currentProfileUsername, username, domain)
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProfileFeature) iAddANewLoginForDomainAndUsername(domain, username string) error {
	p.lastErr = p.cli.AddLogin(p.currentProfileUsername, username, domain, "123")
	if !errors.Is(p.lastErr, model.ErrorLoginAlreadyExistsInProfile) {
		return p.lastErr
	}

	return nil
}

func (p *ProfileFeature) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have empty storage$`, p.iHaveEmptyStorage)

	ctx.Step(`^I have a profile for username "([^"]*)"$`, p.iHaveAProfileForUsername)
	ctx.Step(`^I signup using username "([^"]*)"$`, p.iSignupUsingUsername)
	ctx.Step(`^I will get this error "([^"]*)"$`, p.iWillGetThisError)

	ctx.Step(`^I have no profiles stored$`, p.iHaveEmptyStorage)
	ctx.Step(`^I will have a stored profile with username "([^"]*)" and empty list of logins$`, p.iWillHaveAStoredProfileWithUsernameAndEmptyListOfLogins)

	ctx.Step(`^I add a new login for domain "([^"]*)" and username "([^"]*)"$`, p.iAddANewLoginForDomainAndUsername)
	ctx.Step(`^It has a login for domain "([^"]*)" and username "([^"]*)"$`, p.itHasALoginForDomainAndUsername)
}

func NewProfile(cli *presentation.Cli, storageUrl string) *ProfileFeature {
	return &ProfileFeature{cli, storageUrl, "", nil}
}

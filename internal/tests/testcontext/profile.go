package testcontext

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MAZEN-Kenjrawi/pwd/internal/infrastructure"
	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"github.com/MAZEN-Kenjrawi/pwd/internal/presentation"
)

type Profile struct {
	cmdBus     infrastructure.CmdBus
	storageUrl string
	lastErr    error
}

func (p *Profile) IHaveAProfileForUsername(username string) error {
	path := filepath.Join(p.storageUrl, fmt.Sprintf("%s.json", username))
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write([]byte(`{}`))

	return err
}

func (p *Profile) ISignupUsingUsername(username string) error {
	cli := presentation.NewCli(p.cmdBus)
	p.lastErr = cli.Signup(username, "foo-bar")
	if errors.Is(p.lastErr, model.ErrorProfileAlreadyExists) {
		return nil
	}

	return p.lastErr
}

func (p *Profile) IWillGetThisError(expectedError string) error {
	if p.lastErr.Error() == expectedError {
		return nil
	}

	return fmt.Errorf("unexpected error return, expected: `%s`, actual: `%s`", expectedError, p.lastErr)
}

func (p *Profile) IHaveNoProfilesStored() error {
	return p.ClearStorage()
}

func (p *Profile) IWillHaveAStoredProfileWithUsernameAndEmptyListOfLogins(username string) error {
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

func (p *Profile) ClearStorage() error {
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

func NewProfile(cmdBus infrastructure.CmdBus, storageUrl string) *Profile {
	return &Profile{cmdBus, storageUrl, nil}
}

package infrastructure

import (
	"fmt"
	"reflect"

	"github.com/MAZEN-Kenjrawi/pwd/internal/application"
	"github.com/MAZEN-Kenjrawi/pwd/internal/infrastructure/json"
	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"go.uber.org/dig"
)

const FILE_SYSTEM_STORAGE_MODE = "filesystem"

type Container struct {
	*dig.Container
}

func NewContainer(cfg Config) (*Container, error) {
	c := dig.New()
	switch cfg.Storage.Mode {
	case FILE_SYSTEM_STORAGE_MODE:
		if err := c.Provide(func() model.ProfileRepository {
			return json.NewProfileRepository(cfg.Storage.Url)
		}); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("storage mode %s is not supported", cfg.Storage.Mode)
	}

	if err := c.Provide(provideCmdBus); err != nil {
		return nil, err
	}

	if err := c.Provide(provideQueryBus); err != nil {
		return nil, err
	}

	return &Container{c}, nil
}

func provideProfileRepository(filePath string) model.ProfileRepository {
	return json.NewProfileRepository(filePath)
}

func provideCmdBus(profileRepo model.ProfileRepository) CmdBus {
	return CmdBus{
		application.AddLoginHandler{ProfileRepository: profileRepo},
		application.RemoveLoginHandler{ProfileRepository: profileRepo},
		application.SignUpHandler{ProfileRepository: profileRepo},
	}
}

func provideQueryBus(profileRepo model.ProfileRepository) QueryBus {
	return QueryBus{
		application.GetLoginHandler{ProfileRepository: profileRepo},
		application.SearchLoginHandler{ProfileRepository: profileRepo},
	}
}

func getType(v interface{}) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

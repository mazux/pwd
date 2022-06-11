package application

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
)

type AddLoginCommand struct {
	ProfileUsername string
	Username        string
	Domain          string
	Password        string
}

type AddLoginHandler struct {
	ProfileRepository model.ProfileRepository
}

func (h *AddLoginHandler) Handle(cmd AddLoginCommand) error {
	profile, err := h.ProfileRepository.GetProfileByUsername(cmd.ProfileUsername)
	if err != nil {
		return err
	}

	if profile == nil {
		return model.ErrorProfileDoesNotExist
	}

	err = profile.AddLogin(cmd.Username, cmd.Domain, cmd.Password)
	if err != nil {
		return err
	}

	return h.ProfileRepository.Save(profile)
}

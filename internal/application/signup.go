package application

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
)

type SignUpCommand struct {
	Username string
	Secret   string
}

type SignUpHandler struct {
	ProfileRepository model.ProfileRepository
}

func (h *SignUpHandler) Handle(cmd SignUpCommand) error {
	existingProfile, err := h.ProfileRepository.GetProfileByUsername(cmd.Username)
	if err != nil {
		return err
	}

	if existingProfile != nil {
		return model.ErrorProfileAlreadyExists
	}

	profile, err := model.NewProfile(cmd.Username, cmd.Secret)
	if err != nil {
		return err
	}

	return h.ProfileRepository.Save(profile)
}

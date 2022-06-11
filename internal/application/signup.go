package application

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
)

type SignUpCommand struct {
	Username string
	Secret   string
}

type SignUpHandler struct {
	ProfileRepo model.ProfileRepository
}

func (h *SignUpHandler) Handle(cmd SignUpCommand) error {
	existingProfile, err := h.ProfileRepo.GetProfileByUsername(cmd.Username)
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

	return h.ProfileRepo.Save(profile)
}

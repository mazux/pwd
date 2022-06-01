package application

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
)

var ()

type CreateProfileCommand struct {
	Username string
	Secret   string
}

type CreateProfileHandler struct {
	ProfileRepo model.ProfileRepository
}

func (h *CreateProfileHandler) Handle(cmd CreateProfileCommand) error {
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

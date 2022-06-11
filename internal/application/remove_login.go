package application

import "github.com/MAZEN-Kenjrawi/pwd/internal/model"

type RemoveLoginCommand struct {
	ProfileUsername string
	Username        string
	Domain          string
}

type RemoveLoginHandler struct {
	ProfileRepository model.ProfileRepository
}

func (h *RemoveLoginHandler) Handle(cmd RemoveLoginCommand) error {
	profile, err := h.ProfileRepository.GetProfileByUsername(cmd.ProfileUsername)
	if err != nil {
		return err
	}

	if profile == nil {
		return model.ErrorProfileDoesNotExist
	}

	return profile.RemoveLogin(cmd.Username, cmd.Domain)
}

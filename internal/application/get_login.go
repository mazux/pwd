package application

import "github.com/MAZEN-Kenjrawi/pwd/internal/model"

type Login struct {
	Domain   string
	Username string
	Password string
}

type GetLoginQuery struct {
	ProfileUsername string
	Domain          string
	Username        string
}

type GetLoginHandler struct {
	ProfileRepository model.ProfileRepository
}

func (h *GetLoginHandler) Handle(qry GetLoginQuery) (*Login, error) {
	profile, err := h.ProfileRepository.GetProfileByUsername(qry.ProfileUsername)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, model.ErrorProfileDoesNotExist
	}

	login := profile.GetLogins().Filter(func(l *model.Login) bool {
		return l.Domain == qry.Domain && l.Username == qry.Username
	}).First()

	if login == nil {
		return nil, nil
	}

	decryptedPassword, err := login.DecryptPassword()
	if err != nil {
		return nil, err
	}
	return &Login{login.Domain, login.Username, decryptedPassword}, nil
}

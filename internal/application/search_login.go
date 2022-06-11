package application

import "github.com/MAZEN-Kenjrawi/pwd/internal/model"

type LoginList []struct {
	Username string
	Domain   string
}

type SearchLoginQuery struct {
	ProfileUsername string
	Domain          string
}

type SearchLoginHandler struct {
	ProfileRepo model.ProfileRepository
}

func (h *SearchLoginHandler) Handle(qry SearchLoginQuery) (LoginList, error) {
	profile, err := h.ProfileRepo.GetProfileByUsername(qry.ProfileUsername)
	if err != nil {
		return nil, err
	}

	if profile == nil {
		return nil, model.ErrorProfileDoesNotExist
	}

	list := profile.GetLogins().Filter(func(l *model.Login) bool {
		return l.Domain == qry.Domain
	})

	filteredList := make(LoginList, 0)
	for _, l := range list {
		filteredList = append(filteredList, struct {
			Username string
			Domain   string
		}{
			Username: l.Username,
			Domain:   l.Domain,
		})
	}

	return filteredList, nil
}

package tests

import (
	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
)

type ProfileRepository struct {
	Profiles map[string]model.Profile
}

func (r *ProfileRepository) Save(profile *model.Profile) error {
	r.Profiles[profile.Username] = *profile

	return nil
}

func (r *ProfileRepository) GetProfileByUsername(username string) (*model.Profile, error) {
	if p, exist := r.Profiles[username]; exist {
		return &p, nil
	}

	return nil, nil
}

func (r *ProfileRepository) AssertHasProfile(username string) bool {
	_, exist := r.Profiles[username]

	return exist
}

func (r *ProfileRepository) AssertHasProfileWithLogin(profileUsername, loginUsername, domain string) bool {
	p, exist := r.Profiles[profileUsername]
	if !exist {
		return false
	}

	return p.HasLogin(loginUsername, domain)
}

func NewMockProfile(usernam, secret string, loginList []map[string]string) model.Profile {
	p, _ := model.NewProfile(usernam, secret)
	for _, login := range loginList {
		p.AddLogin(login["username"], login["domain"], login["password"])
	}

	return *p
}

func NewMockProfileRepository(profiles ...model.Profile) *ProfileRepository {
	r := &ProfileRepository{make(map[string]model.Profile)}
	for _, p := range profiles {
		r.Profiles[p.Username] = p
	}
	
	return r
}

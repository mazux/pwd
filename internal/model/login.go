package model

import "fmt"

type Login struct {
	Username string
	Domain   string
	password password
}

type loginList []*Login

func (ll loginList) Filter(f func(*Login) bool) loginList {
	var clone loginList
	for _, l := range ll {
		if f(l) {
			clone = append(clone, l)
		}
	}

	return clone
}

func newLogin(username, domain, password string) (*Login, error) {
	key := keyFromString(fmt.Sprintf("%s%s", username, domain))
	encryptedPassword, err := encrypt(password, key)
	if err != nil {
		return nil, err
	}

	return &Login{
		username,
		domain,
		encryptedPassword,
	}, nil
}

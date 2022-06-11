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

func (ll loginList) First() *Login {
	if len(ll) == 0 {
		return nil
	}
	return ll[0]
}

func (l Login) DecryptPassword() (string, error) {
	return l.password.decrypt(l.getKey())
}

func (l *Login) encryptPassword(password string) error {
	encryptedPassword, err := encrypt(password, l.getKey())
	if err != nil {
		return err
	}

	l.password = encryptedPassword
	return nil
}

func (l Login) getKey() key {
	return keyFromString(fmt.Sprintf("%s%s", l.Username, l.Domain))
}

func newLogin(username, domain, password string) (*Login, error) {
	login := &Login{
		Username: username,
		Domain:   domain,
	}

	err := login.encryptPassword(password)
	if err != nil {
		return nil, err
	}

	return login, nil
}

package application

import (
	"errors"
	"testing"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"github.com/MAZEN-Kenjrawi/pwd/internal/tests"
)

func TestGetLoginFromNonExisitingProfile(t *testing.T) {
	// Arrange
	t.Parallel()
	repo := tests.NewMockProfileRepository()
	handler := &GetLoginHandler{repo}
	qry := GetLoginQuery{"Mazen", "foo.com"}

	// Act
	login, err := handler.Handle(qry)

	// Assert
	if !errors.Is(err, model.ErrorProfileDoesNotExist) {
		t.Errorf("unexpected error happen: %s", err)
	}

	if login != nil {
		t.Errorf("login must be empty, got: %v", login)
	}
}

func TestGetLoginFroProfileWithNoMatch(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	loginDomain := "github.com"
	mockProfile := tests.NewMockProfile(profileUsername, "foo", []map[string]string{
		{"username": "MAZux", "domain": loginDomain, "password": "123"},
	})
	repo := tests.NewMockProfileRepository(mockProfile)

	handler := &GetLoginHandler{repo}
	qry := GetLoginQuery{profileUsername, "foo.com"}

	// Act
	login, err := handler.Handle(qry)

	// Assert
	if err != nil {
		t.Errorf("unexpected error happen: %s", err)
	}

	if login != nil {
		t.Errorf("login must be empty, got: %v", login)
	}
}

func TestGetLoginFroProfileWithMatch(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	loginDomain := "github.com"
	mockProfile := tests.NewMockProfile(profileUsername, "foo", []map[string]string{
		{"username": "MAZux", "domain": loginDomain, "password": "123"},
	})
	repo := tests.NewMockProfileRepository(mockProfile)

	handler := &GetLoginHandler{repo}
	qry := GetLoginQuery{profileUsername, loginDomain}

	// Act
	login, err := handler.Handle(qry)

	// Assert
	if err != nil {
		t.Errorf("unexpected error happen: %s", err)
	}

	if login == nil {
		t.Fatalf("login must not be empty")
	}

	if login.Domain != loginDomain {
		t.Errorf("login domain name mismatch, expected %s but got %s", loginDomain, login.Domain)
	}

	if login.Username != "MAZux" {
		t.Errorf("login username mismatch, expected MAZux but got %s", login.Username)
	}

	if login.Password != "123" {
		t.Errorf("login password mismatch, expected 123 but got %s", login.Password)
	}
}

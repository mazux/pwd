package application

import (
	"errors"
	"testing"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"github.com/MAZEN-Kenjrawi/pwd/internal/tests"
)

func TestAddLoginForNotExistingProfile(t *testing.T) {
	// Arrange
	t.Parallel()
	handler := &AddLoginHandler{tests.NewMockProfileRepository()}
	cmd := AddLoginCommand{"Mazen", "MAZux", "stackoverflow.com", "123"}

	// Act
	err := handler.Handle(cmd)

	// Assert
	if !errors.Is(err, model.ErrorProfileDoesNotExist) {
		t.Errorf("unexpected error happen: %s", err)
	}
}

func TestAddDuplicatedLoginToProfile(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	duplicatedLoginUsername := "MAZux"
	duplicatedLoginDomain := "stackoverflow.com"
	mockProfile := tests.NewMockProfile(profileUsername, "foo", []map[string]string{
		{"username": duplicatedLoginUsername, "domain": duplicatedLoginDomain, "password": "123"},
	})
	repo := tests.NewMockProfileRepository(mockProfile)

	handler := &AddLoginHandler{repo}
	cmd := AddLoginCommand{profileUsername, duplicatedLoginUsername, duplicatedLoginDomain, "123"}

	// Act
	err := handler.Handle(cmd)

	// Assert
	if !errors.Is(err, model.ErrorLoginAlreadyExistsInProfile) {
		t.Errorf("unexpected error happen: %s", err)
	}
}

func TestAddLoginToProfileWithNoLogins(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	loginUsername := "MAZux"
	domain := "stackoverflow.com"
	p := tests.NewMockProfile(profileUsername, "foo", []map[string]string{})
	repo := tests.NewMockProfileRepository(p)

	handler := &AddLoginHandler{repo}
	cmd := AddLoginCommand{profileUsername, loginUsername, domain, "123"}

	// Act
	err := handler.Handle(cmd)
	if err != nil {
		t.Errorf("unexpected error happen: %s", err)
	}

	// Assert
	if !repo.AssertHasProfileWithLogin(profileUsername, loginUsername, domain) {
		t.Errorf("login %s was not added to profile %s", loginUsername, profileUsername)
	}
}

func TestAddLoginToProfileWithOneLogin(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	loginUsername := "MAZux"
	domain := "stackoverflow.com"
	mockProfile := tests.NewMockProfile(profileUsername, "foo", []map[string]string{
		{"username": "just@domain.com", "domain": "foo-domain.com", "password": "123"},
	})
	repo := tests.NewMockProfileRepository(mockProfile)
	
	handler := &AddLoginHandler{repo}
	cmd := AddLoginCommand{profileUsername, loginUsername, domain, "123"}

	// Act
	err := handler.Handle(cmd)
	if err != nil {
		t.Errorf("unexpected error happen: %s", err)
	}

	// Assert
	if !repo.AssertHasProfileWithLogin(profileUsername, loginUsername, domain) {
		t.Errorf("login %s was not added to profile %s", loginUsername, profileUsername)
	}
}

package application

import (
	"errors"
	"testing"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"github.com/MAZEN-Kenjrawi/pwd/internal/tests"
)

func TestRemoveLoginForNonExistedProfile(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	repo := tests.NewMockProfileRepository()

	handler := &RemoveLoginHandler{repo}
	cmd := RemoveLoginCommand{profileUsername, "MAZux", "foo"}

	// Act
	err := handler.Handle(cmd)

	// Assert
	if !errors.Is(err, model.ErrorProfileDoesNotExist) {
		t.Errorf("expected error '%s', got '%s'", model.ErrorProfileDoesNotExist, err)
	}
}

func TestRemoveNonExistedLoginFromProfile(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	mockProfile := tests.NewMockProfile(profileUsername, "foo", []map[string]string{
		{"username": "MAZux", "domain": "github.com", "password": "123"},
	})
	repo := tests.NewMockProfileRepository(mockProfile)

	handler := &RemoveLoginHandler{repo}
	cmd := RemoveLoginCommand{profileUsername, "MAZux", "foo"}

	// Act
	err := handler.Handle(cmd)

	// Assert
	if !errors.Is(err, model.ErrorLoginDoesNotExistInProfile) {
		t.Errorf("expected error '%s', got '%s'", model.ErrorLoginDoesNotExistInProfile, err)
	}
}

func TestRemoveLoginFromProfile(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	loginUsername := "MAZux"
	domain := "github.com"
	mockProfile := tests.NewMockProfile(profileUsername, "foo", []map[string]string{
		{"username": loginUsername, "domain": domain, "password": "123"},
	})
	repo := tests.NewMockProfileRepository(mockProfile)

	handler := &RemoveLoginHandler{repo}
	cmd := RemoveLoginCommand{profileUsername, loginUsername, domain}

	// Act
	err := handler.Handle(cmd)

	// Assert
	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}
}

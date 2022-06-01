package application

import (
	"errors"
	"testing"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"github.com/MAZEN-Kenjrawi/pwd/internal/tests"
)

func TestCreateNewProfileForExistingUsername(t *testing.T) {
	// Arrange
	t.Parallel()
	p := tests.NewMockProfile("Mazen", "foo", []map[string]string{})
	repo := tests.NewMockProfileRepository(p)

	cmd := CreateProfileCommand{"Mazen", "foo"}
	handler := &CreateProfileHandler{repo}

	// Act
	err := handler.Handle(cmd)

	// Assert
	if !errors.Is(err, model.ErrorProfileAlreadyExists) {
		t.Errorf("unexpected error happened: %s", err)
	}
}

func TestCreateNewProfileWithNoErrors(t *testing.T) {
	// Arrange
	t.Parallel()
	username := "Mazen"
	repo := tests.NewMockProfileRepository()
	
	handler := &CreateProfileHandler{repo}
	cmd := CreateProfileCommand{username, "foo"}

	// Act
	err := handler.Handle(cmd)
	if err != nil {
		t.Errorf("unexpected error happened: %s", err)
	}

	// Assert
	if !repo.AssertHasProfile(username) {
		t.Errorf("expected profile was not persisted")
	}
}

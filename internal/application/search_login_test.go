package application

import (
	"errors"
	"testing"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
	"github.com/MAZEN-Kenjrawi/pwd/internal/tests"
)

func TestSearchForLoginInNonExistedProfile(t *testing.T) {
	// Arrange
	t.Parallel()
	handler := &SearchLoginHandler{tests.NewMockProfileRepository()}
	qry := SearchLoginQuery{"Mazen", "foo"}

	// Act
	list, err := handler.Handle(qry)
	// Assert
	if !errors.Is(err, model.ErrorProfileDoesNotExist) {
		t.Errorf("unexpected error '%s'", err)
	}

	if len(list) != 0 {
		t.Errorf("search list must be empty, instead it has %d items", len(list))
	}
}

func TestSearchForLoginInEmptyProfile(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	p := tests.NewMockProfile(profileUsername, "foo", []map[string]string{})
	repo := tests.NewMockProfileRepository(p)

	handler := &SearchLoginHandler{repo}
	qry := SearchLoginQuery{profileUsername, "foo"}

	// Act
	list, err := handler.Handle(qry)

	// Assert
	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if len(list) != 0 {
		t.Errorf("search list must be empty, instead it has %d items", len(list))
	}
}

func TestSearchForLoginInProfileByOnlyProfileUsername(t *testing.T) {
	// Arrange
	t.Parallel()
	profileUsername := "Mazen"
	domain := "stackoverflow.com"
	username := "MAZux2"
	mockProfile := tests.NewMockProfile(profileUsername, "foo", []map[string]string{
		{"username": "MAZux", "domain": "github.com", "password": "123"},
		{"username": username, "domain": domain, "password": "432"},
	})
	repo := tests.NewMockProfileRepository(mockProfile)

	handler := &SearchLoginHandler{repo}
	qry := SearchLoginQuery{profileUsername, domain}

	// Act
	list, err := handler.Handle(qry)

	// Assert
	if err != nil {
		t.Errorf("unexpected error '%s'", err)
	}

	if len(list) == 0 {
		t.Fatalf("search list must not be empty")
	}

	searchItem := list[0]
	if searchItem.Domain != domain {
		t.Errorf("domain name of search result is not correct, expected: %s, got: %s", domain, searchItem.Domain)
	}

	if searchItem.Username != username {
		t.Errorf("domain name of search result is not correct, expected: %s, got: %s", username, searchItem.Username)
	}
}

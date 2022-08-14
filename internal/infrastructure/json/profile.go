package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MAZEN-Kenjrawi/pwd/internal/model"
)

type ProfileRepository struct {
	filePath string
}

func (r *ProfileRepository) Save(p *model.Profile) error {
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", p.Username)
	return os.WriteFile(filepath.Join(r.filePath, filename), j, 0644)
}

func (r *ProfileRepository) GetProfileByUsername(username string) (*model.Profile, error) {
	filename := fmt.Sprintf("%s.json", username)
	file, err := os.Open(filepath.Join(r.filePath, filename))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	profile := model.Profile{}
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func NewProfileRepository(filePath string) model.ProfileRepository {
	return &ProfileRepository{filePath}
}

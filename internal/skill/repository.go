package skill

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Repository interface {
	Exists(name string) (bool, error)
	Extract(name string) (*Skill, error)
	Delete(name string) error
	List() ([]string, error)
	Create(name string) error
}

type skillRepository struct {
	root string
}

func NewSkillRepository(root string) Repository {
	return &skillRepository{root: root}
}

func (cr *skillRepository) Exists(name string) (bool, error) {
	name = name + ".json"
	entries, err := os.ReadDir(cr.root)
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		if entry.Type().IsRegular() && entry.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

func (cr *skillRepository) Extract(name string) (*Skill, error) {
	name = name + ".json"
	skillPath := filepath.Join(cr.root, name)
	if _, err := os.Stat(skillPath); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(skillPath)
	if err != nil {
		return nil, err
	}

	var skill Skill

	if err = json.Unmarshal(data, &skill); err != nil {
		return nil, err
	}

	return &skill, nil

}

func (cr *skillRepository) Delete(name string) error {
	name = name + ".json"
	skillPath := filepath.Join(cr.root, name)
	return os.RemoveAll(skillPath)
}

func (cr *skillRepository) List() ([]string, error) {
	entries, err := os.ReadDir(cr.root)
	if err != nil {
		return nil, err
	}

	results := []string{}
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			results = append(results, entry.Name())
		}
	}

	return results, nil

}

func (cr *skillRepository) Create(name string) error {

	skill := Skill{Metadata: Metadata{Name: name}}

	name = name + ".json"
	skillPath := filepath.Join(cr.root, name)

	data, err := json.MarshalIndent(skill, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(skillPath, data, 0644)
}

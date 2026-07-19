package workspace

import (
	"os"
	"path/filepath"
)

type Repository interface {
	Exists(name string) (bool, error)
	Create(name string) error
	Delete(name string) error
	List() ([]string, error)
}

type workspaceRepository struct {
	root string
}

func NewWorkspaceRepository(root string) Repository {
	return &workspaceRepository{root: root}
}

func (wr *workspaceRepository) Exists(name string) (bool, error) {
	entries, err := os.ReadDir(wr.root)
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

func (wr *workspaceRepository) Create(name string) error {
	path := filepath.Join(wr.root, name)
	return os.MkdirAll(path, 0755)
}

func (wr *workspaceRepository) Delete(name string) error {
	path := filepath.Join(wr.root, name)
	return os.RemoveAll(path)
}

func (wr *workspaceRepository) List() ([]string, error) {
	results := []string{}
	entries, err := os.ReadDir(wr.root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			results = append(results, entry.Name())
		}
	}

	return results, nil

}

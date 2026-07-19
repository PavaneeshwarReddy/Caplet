package workspace

import (
	"caplet/internal/config"
	"log"
)

type Service struct {
	workspaceRepo Repository
	configRepo    config.Repository
}

func NewService(workspaceRepo Repository, configRepo config.Repository) *Service {
	return &Service{workspaceRepo: workspaceRepo, configRepo: configRepo}
}

func (s *Service) CreateWorkspace(name string) error {
	exists, err := s.workspaceRepo.Exists(name)
	if err != nil {
		return err
	}
	if exists {
		return ErrWorkspaceAlreadyExists
	}

	return s.workspaceRepo.Create(name)
}

func (s *Service) DeleteWorkspace(name string) error {
	exists, err := s.workspaceRepo.Exists(name)
	if err != nil {
		return err
	}
	if !exists {
		return ErrWorkspaceNotFound
	}

	return s.workspaceRepo.Delete(name)
}

func (s *Service) ListWorkspaces() error {
	entries, err := s.workspaceRepo.List()
	if err != nil {
		return err
	}
	for _, entry := range entries {
		log.Println(entry)
	}
	return nil
}

func (s *Service) UseWorkspace(name string) error {
	exists, err := s.workspaceRepo.Exists(name)
	if err != nil {
		return err
	}
	if !exists {
		return ErrWorkspaceNotFound
	}

	cfg, err := s.configRepo.Get()
	if err != nil {
		return err
	}

	cfg.ActiveWorkspace = name

	return s.configRepo.Save(cfg)
}

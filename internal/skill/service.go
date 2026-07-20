package skill

import "log"

type Service struct {
	repo Repository
}

func NewSkillService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateSkill(name string) error {
	exists, err := s.repo.Exists(name)
	if err != nil {
		return err
	}

	if exists {
		return ErrSkillAlreadyExists
	}

	return s.repo.Create(name)
}

func (s *Service) DeleteSkill(name string) error {
	exists, err := s.repo.Exists(name)
	if err != nil {
		return err
	}

	if !exists {
		return ErrSkillNotFound
	}

	return s.repo.Delete(name)

}

func (s *Service) ListSkills() error {
	skills, err := s.repo.List()
	if err != nil {
		return err
	}

	for _, skill := range skills {
		log.Println(skill)
	}

	return nil
}

func (s *Service) InspectSkill(name string) error {
	exists, err := s.repo.Exists(name)
	if err != nil {
		return err
	}

	if !exists {
		return ErrSkillNotFound
	}

	skill, err := s.repo.Extract(name)
	if err != nil {
		return err
	}
	log.Println(skill)
	return nil
}

func (s *Service) SyncSkills() error {
	return nil
}

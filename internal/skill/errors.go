package skill

import "errors"

var (
	ErrSkillNotFound      = errors.New("Skill not found")
	ErrSkillAlreadyExists = errors.New("Skill already exists")
)

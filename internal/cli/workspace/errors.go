package workspace

import "errors"

var (
	ErrWorkspaceNotFound      = errors.New("workspace not found")
	ErrWorkspaceAlreadyExists = errors.New("workspace already exists")
	ErrInvalidWorkspaceName   = errors.New("invalid workspace name")
	ErrNoActiveWorkspace      = errors.New("no active workspace")
)

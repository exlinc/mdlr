package vcs

import "errors"

var (
	ErrInvalidVCSType    = errors.New("error invalid VCS type")
	ErrGitNotAvailable   = errors.New("error git command not found. please make sure that git is installed and available in the PATH")
	ErrRootAlreadyExists = errors.New("error root already exists, try forcing to delete and import")
	ErrRootNotExist      = errors.New("error root does not exist, try importing first")
)

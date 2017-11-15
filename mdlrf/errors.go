package mdlrf

import "errors"

var (
	ErrMdlrFileNotExist        = errors.New("error mdlr.yml file does not exist")
	ErrMdlrFileInvalidPath     = errors.New("error invalid path to mdl.yml file")
	ErrMdlrFileAlreadyLoaded   = errors.New("error mdlr.yml file already loaded")
	ErrMdlrFileAlreadyExists   = errors.New("error mdlr.yml file already exists")
	ErrInvalidSyntaxInMdlrFile = errors.New("error invalid syntax tag in mdlr.yml file. valid syntax tags include: 1")
	ErrInvalidMdlrFile         = errors.New("error invalid mdlr.yml file")
	ErrInvalidModuleDefinition = errors.New("error invalid module definition")
	ErrInvalidModuleType       = errors.New("error invalid module type. must be either git or hg")
)

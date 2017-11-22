package mdlrf

import (
	"git.exlhub.io/exlinc/tools-mdlr/config"
	"git.exlhub.io/exlinc/tools-mdlr/vcs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Log = config.Cfg().GetLogger()

var mdlrFileHeader = []byte(`# Welcome to mdlr.yml
# Use the mdlr CLI to edit and explore this file
# WARNING: The mdlr CLI does not preserve comments!

`)

type MdlrFile struct {
	ParentDirectory  string             `yaml:"-"`
	AbsoluteFilePath string             `yaml:"-"`
	Syntax           int64              `yaml:"syntax"`
	Modules          map[string]*Module `yaml:"modules,flow"`
}

func NewMdlrFile() *MdlrFile {
	return &MdlrFile{
		Syntax:  1,
		Modules: map[string]*Module{},
	}
}

type Module struct {
	Name         string `yaml:"-"` // This is the key used for the module -- populated outside of the YAML object
	Type         string `yaml:"type"`
	Path         string `yaml:"path"`
	AbsolutePath string `yaml:"-"`
	URL          string `yaml:"url"`
	Branch       string `yaml:"branch"`
	Commit       string `yaml:"commit"`
	vcs.Context  `yaml:"-"`
}

func (mod *Module) Validate() error {
	if mod.Name == "" || mod.Type == "" || mod.Path == "" || mod.URL == "" || mod.AbsolutePath == "" {
		return ErrInvalidModuleDefinition
	}
	if !vcs.Supported(mod.Type) {
		return ErrInvalidModuleType
	}
	var err error
	// TODO: set verbose
	mod.Context, err = vcs.Load(true, mod.Type, mod.AbsolutePath)
	if err != nil {
		return err
	}
	return nil
}

func (mod *Module) Prepare(name string, parentDir string) {
	// TODO - fill defaults
	mod.Name = name
	// Default to git
	if mod.Type == "" {
		mod.Type = "git"
	}
	// Try to set defaults based on the type of the module
	switch mod.Type {
	case "git":
		if mod.Branch == "" {
			mod.Branch = "master"
		}
		if mod.Commit == "" {
			mod.Commit = "HEAD"
		}
	}
	mod.AbsolutePath = mod.Path
	if !filepath.IsAbs(mod.AbsolutePath) {
		mod.AbsolutePath = filepath.Join(parentDir, mod.AbsolutePath)
	}
}

func (mf *MdlrFile) Exists(absFilePath string) bool {
	if _, err := os.Stat(absFilePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func (mf *MdlrFile) Load(absFilePath string) error {
	if !mf.Exists(absFilePath) {
		return ErrMdlrFileNotExist
	}
	var err error
	if !filepath.IsAbs(absFilePath) {
		absFilePath, err = filepath.Abs(absFilePath)
		if err != nil {
			return err
		}
	}
	raw, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, mf)
	if err != nil {
		return err
	}
	mf.Prepare(absFilePath)
	return mf.Validate()
}

func (mf *MdlrFile) Prepare(absFilePath string) {
	mf.AbsoluteFilePath = absFilePath
	mf.ParentDirectory = filepath.Dir(mf.AbsoluteFilePath)
	if mf.Syntax == 0 {
		mf.Syntax = 1
	}
	if mf.Modules == nil {
		mf.Modules = make(map[string]*Module)
	}
	for k, v := range mf.Modules {
		v.Prepare(k, mf.ParentDirectory)
		mf.Modules[k] = v
	}
}

func (mf *MdlrFile) Persist() error {
	out, err := yaml.Marshal(mf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(mf.AbsoluteFilePath, append(mdlrFileHeader, out...), 0644)
}

func (mf *MdlrFile) Validate() error {
	if mf.Syntax != 1 {
		return ErrInvalidSyntaxInMdlrFile
	}
	if mf.Modules == nil {
		return ErrInvalidMdlrFile
	}
	for k := range mf.Modules {
		err := mf.Modules[k].Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

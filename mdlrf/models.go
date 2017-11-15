package mdlrf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var mdlrFileHeader = []byte(`# Welcome to mdlr.yml
# Use the mdlr CLI to edit and explore this file
# WARNING: The mdlr CLI does not preserve comments!

`)

type MdlrFile struct {
	Syntax  int64              `yaml:"syntax"`
	Modules map[string]*Module `yaml:"modules,flow"`
}

func NewMdlrFile() *MdlrFile {
	return &MdlrFile{
		Syntax:  1,
		Modules: map[string]*Module{},
	}
}

type Module struct {
	Name   string `yaml:"-"` // This is the key used for the module -- populated outside of the YAML object
	Type   string `yaml:"type"`
	Path   string `yaml:"path"`
	URL    string `yaml:"url"`
	Branch string `yaml:"branch"`
	Commit string `yaml:"commit"`
}

func (mod *Module) Validate() error {
	if mod.Name == "" || mod.Type == "" || mod.Path == "" || mod.URL == "" || mod.Branch == "" || mod.Commit == "" {
		return ErrInvalidModuleDefinition
	}
	if mod.Type != "git" && mod.Type != "hg" {
		return ErrInvalidModuleType
	}
	return nil
}

func (mf *MdlrFile) Exists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func (mf *MdlrFile) Load(filename string) error {
	if !mf.Exists(filename) {
		return ErrMdlrFileNotExist
	}
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, mf)
	if err != nil {
		return err
	}
	mf.Prepare()
	return nil
}

func (mf *MdlrFile) Prepare() {
	if mf.Syntax == 0 {
		mf.Syntax = 1
	}
	if mf.Modules == nil {
		mf.Modules = make(map[string]*Module)
	}
	for k, v := range mf.Modules {
		v.Name = k
		// Add some sensible defaults for the optional fields
		if v.Branch == "" {
			v.Branch = "master"
		}
		if v.Commit == "" {
			v.Commit = "HEAD"
		}
		if v.Type == "" {
			v.Type = "git"
		}
		mf.Modules[k] = v
	}
}

func (mf *MdlrFile) Persist(filename string) error {
	out, err := yaml.Marshal(mf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, append(mdlrFileHeader, out...), 0644)
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

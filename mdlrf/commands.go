package mdlrf

import (
	"fmt"
	"os"
)

type MdlrCtx struct {
	IsFileReady bool
	FilePath    string
	MdlrFile    *MdlrFile
}

func NewMdlrCtxForCmd() (*MdlrCtx, error) {
	c := &MdlrCtx{}
	var err error
	c.FilePath, err = getMdlrFilePathForCmd()
	return c, err
}

func (ctx *MdlrCtx) loadFile() error {
	if ctx.MdlrFile != nil || ctx.IsFileReady {
		return ErrMdlrFileAlreadyLoaded
	}
	if ctx.FilePath == "" {
		return ErrMdlrFileInvalidPath
	}
	ctx.MdlrFile = &MdlrFile{}
	err := ctx.MdlrFile.Load(ctx.FilePath)
	if err != nil {
		ctx.MdlrFile = nil
		return err
	}
	ctx.IsFileReady = true
	return nil
}

func (ctx *MdlrCtx) Init() error {
	if err := ctx.loadFile(); err == nil {
		return ErrMdlrFileAlreadyExists
	} else if err != ErrMdlrFileNotExist {
		return err
	}
	ctx.MdlrFile = NewMdlrFile()
	ctx.MdlrFile.Prepare(ctx.FilePath)
	ctx.IsFileReady = true
	return ctx.MdlrFile.Persist()
}

func (ctx *MdlrCtx) List() (string, error) {
	// TODO
	err := ctx.loadFile()
	if err != nil {
		return "", err
	}
	if len(ctx.MdlrFile.Modules) == 0 {
		return "There aren't any modules defined in the mdlr.yml file yet. Try running the add command with mdlr to add a module.", nil
	}
	items := make([]string, 0, len(ctx.MdlrFile.Modules))
	for _, m := range ctx.MdlrFile.Modules {
		items = append(items, fmt.Sprintf("[%s] %s -> %s (%s) hosted at %s on branch %s at %s", m.Status(true), m.Path, m.Name, m.Type, m.URL, m.Branch, m.Commit))
	}
	out := fmt.Sprintf("Modules count: %d", len(items))
	for _, val := range items {
		out += fmt.Sprintf("\n\t%s", val)
	}
	return out, nil
}

func (ctx *MdlrCtx) Add(name string, mType string, path string, url string, branch string, commit string) error {
	err := ctx.loadFile()
	if err != nil {
		return err
	}
	if _, exist := ctx.MdlrFile.Modules[name]; exist {
		return ErrModuleNameAlreadyInUse
	}
	ctx.MdlrFile.Modules[name] = &Module{
		Type:   mType,
		Path:   path,
		URL:    url,
		Branch: branch,
		Commit: commit,
	}
	ctx.MdlrFile.Modules[name].Prepare(name, ctx.MdlrFile.ParentDirectory)
	err = ctx.MdlrFile.Modules[name].Validate()
	if err != nil {
		return err
	}
	return ctx.MdlrFile.Persist()
}

func (ctx *MdlrCtx) Remove(name string, dropFiles bool) error {
	err := ctx.loadFile()
	if err != nil {
		return err
	}
	if len(ctx.MdlrFile.Modules) == 0 {
		return ErrNoModules
	}
	if _, exist := ctx.MdlrFile.Modules[name]; !exist {
		return ErrModuleNameNotExist
	}
	dirPath := ctx.MdlrFile.Modules[name].AbsolutePath
	delete(ctx.MdlrFile.Modules, name)
	if dropFiles {
		return os.RemoveAll(dirPath)
	}
	return nil
}

func (ctx *MdlrCtx) Import(specificName string, force bool) error {
	// TODO
	err := ctx.loadFile()
	if err != nil {
		return err
	}
	if len(ctx.MdlrFile.Modules) == 0 {
		return ErrNoModules
	}
	switch specificName {
	case "":
		// TODO
	default:
		if _, exist := ctx.MdlrFile.Modules[specificName]; !exist {
			return ErrModuleNameNotExist
		}
		// TODO
	}
	return nil
}

func (ctx *MdlrCtx) Update(specificName string, force bool) error {
	// TODO
	err := ctx.loadFile()
	if err != nil {
		return err
	}
	if len(ctx.MdlrFile.Modules) == 0 {
		return ErrNoModules
	}
	switch specificName {
	case "":
		// TODO
	default:
		if _, exist := ctx.MdlrFile.Modules[specificName]; !exist {
			return ErrModuleNameNotExist
		}
		// TODO
	}
	return nil
}

func (ctx *MdlrCtx) Status(name string) (string, error) {
	err := ctx.loadFile()
	if err != nil {
		return "", err
	}
	if len(ctx.MdlrFile.Modules) == 0 {
		return "", ErrNoModules
	}
	if _, exist := ctx.MdlrFile.Modules[name]; !exist {
		return "", ErrModuleNameNotExist
	}
	return ctx.MdlrFile.Modules[name].Status(false), nil
}

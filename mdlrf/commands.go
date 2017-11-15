package mdlrf

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
	ctx.IsFileReady = true
	ctx.MdlrFile.Persist(ctx.FilePath)
	return nil
}

func (ctx *MdlrCtx) List() error {
	// TODO
	return nil
}

func (ctx *MdlrCtx) Add() error {
	// TODO
	return nil
}

func (ctx *MdlrCtx) Remove() error {
	// TODO
	return nil
}

func (ctx *MdlrCtx) Import() error {
	// TODO
	return nil
}

func (ctx *MdlrCtx) Update() error {
	// TODO
	return nil
}

func (ctx *MdlrCtx) Status() error {
	// TODO
	return nil
}

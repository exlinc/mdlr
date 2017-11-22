package vcs

import (
	"os"
	"path/filepath"
	"strings"
)

type GitVCSCtx struct {
	ParentDir string
	Root      string
	URL       string
	Verbose   bool
}

func setupGitVCSCtx(verbose bool, root string, url string) (Context, error) {
	if !cmdExists("git") {
		return nil, ErrGitNotAvailable
	}
	ctx := &GitVCSCtx{
		ParentDir: filepath.Dir(root),
		Root:      root,
		Verbose:   verbose,
		URL:       url,
	}
	return ctx, nil
}

func (ctx *GitVCSCtx) runCmdInRoot(args ...string) (string, error) {
	return runCmd(ctx.Verbose, ctx.Root, "git", args...)
}

func (ctx *GitVCSCtx) runCmdInParent(args ...string) (string, error) {
	return runCmd(ctx.Verbose, ctx.ParentDir, "git", args...)
}

func (ctx *GitVCSCtx) rootExists() bool {
	if _, err := os.Stat(ctx.Root); os.IsNotExist(err) {
		return false
	}
	return true
}

func (ctx *GitVCSCtx) Import(branch, commit string) error {
	if ctx.rootExists() {
		return ErrRootAlreadyExists
	}
	_, fn := filepath.Split(ctx.Root)
	if _, err := ctx.runCmdInParent("clone", "-b", branch, ctx.URL, fn); err != nil {
		return err
	}
	if _, err := ctx.runCmdInRoot("checkout", commit); err != nil {
		return err
	}
	return nil
}

func (ctx *GitVCSCtx) Update(branch, commit string) error {
	// TODO
	if !ctx.rootExists() {
		return ErrRootNotExist
	}

	return nil
}

func (ctx *GitVCSCtx) Status(short bool) string {
	// TODO - get branch, commit, modified files, etc.
	switch short {
	case true:
		if !ctx.rootExists() {
			return "NONE"
		}
		if out, err := ctx.runCmdInRoot("show", "--oneline", "-s"); err != nil {
			return "ERR"
		} else {
			sep := strings.Split(out, " ")
			if len(sep) > 0 {
				return sep[0]
			} else {
				return "ERR"
			}
		}
	default:
		if !ctx.rootExists() {
			return "The path of the module does not yet exist. Import it to get setup"
		}
		if out, err := ctx.runCmdInRoot("status"); err != nil {
			return "The module is not currently setup or is setup incorrectly. Import it to get setup"
		} else {
			return out
		}
	}
}

func (ctx *GitVCSCtx) Invokable() (bool, error) {
	// TODO
	return false, nil
}

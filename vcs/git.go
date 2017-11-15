package vcs

import (
	"os"
	"path/filepath"
	"strings"
)

type GitVCSCtx struct {
	ParentDir string
	Root      string
	Verbose   bool
}

func setupGitVCSCtx(verbose bool, root string) (Context, error) {
	if !cmdExists("git") {
		return nil, ErrGitNotAvailable
	}
	ctx := &GitVCSCtx{
		ParentDir: filepath.Dir(root),
		Root:      root,
		Verbose:   verbose,
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

func (ctx *GitVCSCtx) Import() error {
	// TODO
	return nil
}

func (ctx *GitVCSCtx) Update() error {
	// TODO
	return nil
}

func (ctx *GitVCSCtx) Status() string {
	// TODO - get branch, commit, modified files, etc.
	if !ctx.rootExists() {
		return "NONE"
	}
	if out, err := ctx.runCmdInRoot("status"); strings.HasPrefix(out, "fatal:") || err != nil {
		return "INVALID"
	}
	return "IMPORTED"
}

func (ctx *GitVCSCtx) Invokable() (bool, error) {
	// TODO
	return false, nil
}

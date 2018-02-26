package vcs

import "exlgit.com/tools/mdlr/config"

var Log = config.Cfg().GetLogger()

type vcsLoadFunc func(verbose bool, root string, url string) (Context, error)

var vcsLoaders = map[string]vcsLoadFunc{
	"git": setupGitVCSCtx,
}

// This doesn't actually check if it's supported on the underlying system -- just with in mdlr
// When the VCS is actually loaded, then it will check whether or not the client is supported on the system
func Supported(vcsType string) bool {
	_, exist := vcsLoaders[vcsType]
	return exist
}

func Load(verbose bool, vcsType string, root string, url string) (Context, error) {
	if !Supported(vcsType) {
		return nil, ErrInvalidVCSType
	}
	return vcsLoaders[vcsType](verbose, root, url)
}

type Context interface {
	Import(branch, commit string) error
	Update(branch, commit string) (string, error)
	Status(short bool) string
	Invokable() (bool, error)
}

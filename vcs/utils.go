package vcs

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func cmdExists(cmdName string) bool {
	_, err := exec.LookPath(cmdName)
	if err != nil {
		return false
	}
	return true
}

type bufferMux struct {
	fwd  bool
	buf  bytes.Buffer
	dest io.Writer
}

func (mux *bufferMux) Write(p []byte) (n int, err error) {
	n, err = mux.buf.Write(p)
	if err != nil || !mux.fwd {
		return
	}
	return mux.dest.Write(p)
}

func (mux *bufferMux) Bytes() []byte {
	return mux.buf.Bytes()
}

func runCmd(verbose bool, wDir string, cmdName string, args ...string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = wDir
	if verbose {
		Log.Infof("cd %s", wDir)
		Log.Infof("%s %s", cmdName, strings.Join(args, " "))
	}
	cmd.Stdin = os.Stdin
	buf := bufferMux{fwd: verbose, buf: bytes.Buffer{}, dest: os.Stdout}
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	out := buf.Bytes()
	if err != nil {
		if verbose {
			Log.Errorf("# cd %s; %s %s", wDir, cmdName, strings.Join(args, " "))
			Log.Error(out)
		}
		return "", err
	}
	return string(out), nil
}

func removeContentsOfDirectory(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

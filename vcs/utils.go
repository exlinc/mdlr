package vcs

import (
	"bytes"
	"fmt"
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

func runCmd(verbose bool, wDir string, cmdName string, args ...string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = wDir
	if verbose {
		Log.Info("cd %s\n", wDir)
		fmt.Printf("%s %s\n", cmdName, strings.Join(args, " "))
	}
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	out := buf.Bytes()
	if err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "# cd %s; %s %s\n", wDir, cmdName, strings.Join(args, " "))
			os.Stderr.Write(out)
		}
		return "", err
	}
	if verbose {
		fmt.Fprintf(os.Stdout, "# cd %s; %s %s\n", wDir, cmdName, strings.Join(args, " "))
		os.Stdout.Write(out)
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

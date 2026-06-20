package commands

import (
	"errors"
	"os"
	"os/exec"
)

func init() {
	Register(&gitCmd{})
}

type gitCmd struct{}

func (g *gitCmd) Name() string        { return "蠢货" }
func (g *gitCmd) Pinyin() string      { return "chǔnhuò" }
func (g *gitCmd) Description() string { return "run git with the given arguments" }
func (g *gitCmd) Execute(args []string) error {
	if _, err := exec.LookPath("git"); err != nil {
		return errors.New("git 未安装  (git is not installed)")
	}
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

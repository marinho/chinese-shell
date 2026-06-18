package commands

import (
	"errors"
	"os"
	"os/exec"
)

func init() {
	Register(&runCmd{})
}

type runCmd struct{}

func (r *runCmd) Name() string        { return "运行" }
func (r *runCmd) Pinyin() string      { return "yùnxíng" }
func (r *runCmd) Description() string { return "run a program at the given path" }
func (r *runCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 运行 <path> [args...]  (usage: 运行 <path> [args...])")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

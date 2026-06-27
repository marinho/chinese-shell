package commands

import (
	"errors"
	"os"
	"os/exec"
)

func init() {
	Register(&lessCmd{})
}

type lessCmd struct{}

func (l *lessCmd) Name() string        { return "少" }
func (l *lessCmd) Pinyin() string      { return "shǎo" }
func (l *lessCmd) Linux() string       { return "less" }
func (l *lessCmd) Description() string { return "view file contents with a pager" }
func (l *lessCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 少 <file>  (usage: 少 <file>)")
	}
	cmd := exec.Command("less", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

package commands

import (
	"errors"
	"os"
	"os/exec"
)

func init() {
	Register(&editCmd{})
}

type editCmd struct{}

func (e *editCmd) Name() string        { return "编辑" }
func (e *editCmd) Pinyin() string      { return "biānjí" }
func (e *editCmd) Linux() string       { return "vim" }
func (e *editCmd) Description() string { return "edit a file in vim" }
func (e *editCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 编辑 <file>  (usage: 编辑 <file>)")
	}
	cmd := exec.Command("vim", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

package commands

import (
	"errors"
	"os"
	"os/exec"
)

func init() {
	Register(&moreCmd{})
}

type moreCmd struct{}

func (m *moreCmd) Name() string        { return "多" }
func (m *moreCmd) Pinyin() string      { return "duō" }
func (m *moreCmd) Linux() string       { return "more" }
func (m *moreCmd) Description() string { return "view file contents with a pager (forward only)" }
func (m *moreCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 多 <file>  (usage: 多 <file>)")
	}
	cmd := exec.Command("more", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

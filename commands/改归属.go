package commands

import (
	"errors"
	"os"
	"os/exec"
)

func init() {
	Register(&chownCmd{})
}

type chownCmd struct{}

func (c *chownCmd) Name() string        { return "改归属" }
func (c *chownCmd) Pinyin() string      { return "gǎi guīshǔ" }
func (c *chownCmd) Description() string { return "change file owner" }
func (c *chownCmd) Execute(args []string) error {
	if len(args) < 2 {
		return errors.New("用法: 改归属 <owner> <file>  (usage: 改归属 <owner> <file>)")
	}
	if _, err := exec.LookPath("chown"); err != nil {
		return errors.New("chown 未安装  (chown is not installed)")
	}
	cmd := exec.Command("chown", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

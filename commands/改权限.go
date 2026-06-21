package commands

import (
	"errors"
	"os"
	"os/exec"
)

func init() {
	Register(&chmodCmd{})
}

type chmodCmd struct{}

func (c *chmodCmd) Name() string        { return "改权限" }
func (c *chmodCmd) Pinyin() string      { return "gǎi quánxiàn" }
func (c *chmodCmd) Description() string { return "change file permissions" }
func (c *chmodCmd) Execute(args []string) error {
	if len(args) < 2 {
		return errors.New("用法: 改权限 <mode> <file>  (usage: 改权限 <mode> <file>)")
	}
	if _, err := exec.LookPath("chmod"); err != nil {
		return errors.New("chmod 未安装  (chmod is not installed)")
	}
	cmd := exec.Command("chmod", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

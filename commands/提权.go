package commands

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func init() {
	Register(&sudoCmd{})
}

type sudoCmd struct{}

func (s *sudoCmd) Name() string        { return "提权" }
func (s *sudoCmd) Pinyin() string      { return "tí quán" }
func (s *sudoCmd) Linux() string       { return "sudo" }
func (s *sudoCmd) Description() string { return "run a command as superuser" }
func (s *sudoCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 提权 <command> [args...]  (usage: 提权 <command> [args...])")
	}
	self, err := os.Executable()
	if err != nil {
		return err
	}
	cmd := exec.Command("sudo", self, "-c", strings.Join(args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

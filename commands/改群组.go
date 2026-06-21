package commands

import (
	"errors"
	"os"
	"os/exec"
)

func init() {
	Register(&chgrpCmd{})
}

type chgrpCmd struct{}

func (c *chgrpCmd) Name() string        { return "改群组" }
func (c *chgrpCmd) Pinyin() string      { return "gǎi qúnzǔ" }
func (c *chgrpCmd) Description() string { return "change file group" }
func (c *chgrpCmd) Execute(args []string) error {
	if len(args) < 2 {
		return errors.New("用法: 改群组 <group> <file>  (usage: 改群组 <group> <file>)")
	}
	if _, err := exec.LookPath("chgrp"); err != nil {
		return errors.New("chgrp 未安装  (chgrp is not installed)")
	}
	cmd := exec.Command("chgrp", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

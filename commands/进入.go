package commands

import (
	"errors"
	"os"
)

func init() {
	Register(&cdCmd{})
}

type cdCmd struct{}

func (c *cdCmd) Name() string        { return "进入" }
func (c *cdCmd) Pinyin() string      { return "jìnrù" }
func (c *cdCmd) Description() string { return "change current directory" }
func (c *cdCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 进入 <path>  (usage: 进入 <path>)")
	}
	return os.Chdir(args[0])
}

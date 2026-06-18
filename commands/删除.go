package commands

import (
	"errors"
	"os"
)

func init() {
	Register(&rmCmd{})
}

type rmCmd struct{}

func (r *rmCmd) Name() string        { return "删除" }
func (r *rmCmd) Pinyin() string      { return "shānchú" }
func (r *rmCmd) Description() string { return "remove a file or directory" }
func (r *rmCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 删除 <path>  (usage: 删除 <path>)")
	}
	for _, path := range args {
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}
	return nil
}

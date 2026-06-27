package commands

import (
	"errors"
	"os"
)

func init() {
	Register(&mvCmd{})
}

type mvCmd struct{}

func (m *mvCmd) Name() string        { return "移动" }
func (m *mvCmd) Pinyin() string      { return "yídòng" }
func (m *mvCmd) Linux() string       { return "mv" }
func (m *mvCmd) Description() string { return "move or rename a file or directory" }
func (m *mvCmd) Execute(args []string) error {
	if len(args) != 2 {
		return errors.New("用法: 移动 <source> <destination>  (usage: 移动 <src> <dst>)")
	}
	return os.Rename(args[0], args[1])
}

package commands

import (
	"errors"
	"os"
)

func init() {
	Register(&setenvCmd{})
}

type setenvCmd struct{}

func (s *setenvCmd) Name() string        { return "设" }
func (s *setenvCmd) Pinyin() string      { return "shè" }
func (s *setenvCmd) Linux() string       { return "export" }
func (s *setenvCmd) Description() string { return "set an environment variable" }
func (s *setenvCmd) Execute(args []string) error {
	if len(args) < 2 {
		return errors.New("用法: 设 <变量> <值>  (usage: 设 <VAR> <value>)")
	}
	return os.Setenv(args[0], args[1])
}

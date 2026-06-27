package commands

import (
	"fmt"
	"strings"
)

func init() {
	Register(&echoCmd{})
}

type echoCmd struct{}

func (e *echoCmd) Name() string        { return "说" }
func (e *echoCmd) Pinyin() string      { return "shuō" }
func (e *echoCmd) Description() string { return "print arguments to stdout" }
func (e *echoCmd) Execute(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

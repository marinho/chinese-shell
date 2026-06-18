package commands

import "os"

func init() {
	Register(&exitCmd{})
}

type exitCmd struct{}

func (e *exitCmd) Name() string        { return "出口" }
func (e *exitCmd) Pinyin() string      { return "chūkǒu" }
func (e *exitCmd) Description() string { return "exit the shell" }
func (e *exitCmd) Execute(_ []string) error {
	println("再见！感谢使用 zhell！")
	println("Goodbye! Thanks for using zhell!")

	os.Exit(0)
	return nil
}

package commands

import (
	"errors"
	"io"
	"os"
)

func init() {
	Register(&catCmd{})
}

type catCmd struct{}

func (c *catCmd) Name() string        { return "猫" }
func (c *catCmd) Pinyin() string      { return "māo" }
func (c *catCmd) Linux() string       { return "cat" }
func (c *catCmd) Description() string { return "print contents of a file" }
func (c *catCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 猫 <file>  (usage: 猫 <file>)")
	}
	for _, path := range args {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(os.Stdout, f)
		f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

package commands

import (
	"errors"
	"fmt"
	"os"
)

func init() {
	Register(&newFileCmd{})
}

type newFileCmd struct{}

func (n *newFileCmd) Name() string        { return "新文件" }
func (n *newFileCmd) Pinyin() string      { return "xīn wénjiàn" }
func (n *newFileCmd) Linux() string       { return "touch" }
func (n *newFileCmd) Description() string { return "create a new empty file" }
func (n *newFileCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 新文件 <file>  (usage: 新文件 <file>)")
	}
	for _, path := range args {
		f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
		if err != nil {
			if os.IsExist(err) {
				return fmt.Errorf("警告: 文件已存在: %s  (file already exists)", path)
			}
			return err
		}
		f.Close()
	}
	return nil
}

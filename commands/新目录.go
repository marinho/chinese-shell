package commands

import (
	"errors"
	"fmt"
	"os"
)

func init() {
	Register(&mkdirCmd{})
}

type mkdirCmd struct{}

func (m *mkdirCmd) Name() string        { return "新目录" }
func (m *mkdirCmd) Pinyin() string      { return "xīn mùlù" }
func (m *mkdirCmd) Linux() string       { return "mkdir" }
func (m *mkdirCmd) Description() string { return "create a new directory" }
func (m *mkdirCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 新目录 <path>  (usage: 新目录 <path>)")
	}
	for _, path := range args {
		info, err := os.Stat(path)
		if err == nil {
			if info.IsDir() {
				fmt.Printf("警告: 目录已存在: %s  (warning: directory already exists)\n", path)
				continue
			}
			return fmt.Errorf("路径已存在且不是目录: %s  (path exists and is not a directory)", path)
		}
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

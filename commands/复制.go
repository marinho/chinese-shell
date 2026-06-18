package commands

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

func init() {
	Register(&cpCmd{})
}

type cpCmd struct{}

func (c *cpCmd) Name() string        { return "复制" }
func (c *cpCmd) Pinyin() string      { return "fùzhì" }
func (c *cpCmd) Description() string { return "copy a file or directory to another path" }
func (c *cpCmd) Execute(args []string) error {
	if len(args) != 2 {
		return errors.New("用法: 复制 <source> <destination>  (usage: 复制 <src> <dst>)")
	}
	return copyPath(args[0], args[1])
}

func copyPath(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return copyDir(src, dst, info)
	}
	return copyFile(src, dst, info.Mode())
}

func copyDir(src, dst string, info os.FileInfo) error {
	if err := os.MkdirAll(dst, info.Mode()); err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if err := copyPath(filepath.Join(src, entry.Name()), filepath.Join(dst, entry.Name())); err != nil {
			return err
		}
	}
	return nil
}

func copyFile(src, dst string, mode os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

package commands

import (
	"fmt"
	"os"
	"os/user"
	"syscall"
)

func init() {
	Register(&statOwnerCmd{})
}

type statOwnerCmd struct{}

func (c *statOwnerCmd) Name() string        { return "看归属" }
func (c *statOwnerCmd) Pinyin() string      { return "kàn guīshǔ" }
func (c *statOwnerCmd) Linux() string       { return "stat" }
func (c *statOwnerCmd) Description() string { return "show owner and group of a file" }
func (c *statOwnerCmd) Execute(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("用法: 看归属 <文件>")
	}
	for _, path := range args {
		info, err := os.Stat(path)
		if err != nil {
			return err
		}
		stat := info.Sys().(*syscall.Stat_t)
		u, err := user.LookupId(fmt.Sprint(stat.Uid))
		if err != nil {
			return err
		}
		g, err := user.LookupGroupId(fmt.Sprint(stat.Gid))
		if err != nil {
			return err
		}
		fmt.Printf("%s  用户: %s  群组: %s\n", path, u.Username, g.Name)
	}
	return nil
}

package commands

import (
	"fmt"
	"os/user"
)

func init() {
	Register(&whoAmICmd{})
}

type whoAmICmd struct{}

func (c *whoAmICmd) Name() string        { return "我是谁" }
func (c *whoAmICmd) Pinyin() string      { return "wǒ shì shuí" }
func (c *whoAmICmd) Linux() string       { return "whoami" }
func (c *whoAmICmd) Description() string { return "print current user" }
func (c *whoAmICmd) Execute(args []string) error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	fmt.Printf("%s (uid: %s)\n", u.Username, u.Uid)
	return nil
}

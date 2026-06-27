package commands

import (
	"fmt"
	"os"
)

func init() {
	Register(&pwdCmd{})
}

type pwdCmd struct{}

func (p *pwdCmd) Name() string        { return "这个目录" }
func (p *pwdCmd) Pinyin() string      { return "zhège mùlù" }
func (p *pwdCmd) Linux() string       { return "pwd" }
func (p *pwdCmd) Description() string { return "print working directory" }
func (p *pwdCmd) Execute(_ []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}

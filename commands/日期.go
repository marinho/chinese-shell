package commands

import (
	"fmt"
	"time"
)

func init() {
	Register(&dateCmd{})
}

type dateCmd struct{}

func (d *dateCmd) Name() string        { return "日期" }
func (d *dateCmd) Pinyin() string      { return "rìqī" }
func (d *dateCmd) Description() string { return "print current date and time" }
func (d *dateCmd) Execute(_ []string) error {
	fmt.Println(time.Now().Format("2006年01月02日 15:04:05 MST"))
	return nil
}

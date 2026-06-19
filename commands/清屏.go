package commands

import (
	"fmt"
)

func init() {
	Register(&clearCmd{})
}

type clearCmd struct{}

func (c *clearCmd) Name() string        { return "清屏" }
func (c *clearCmd) Pinyin() string      { return "qīng píng" }
func (c *clearCmd) Description() string { return "clear the screen" }
func (c *clearCmd) Execute(_ []string) error {
	fmt.Print("\033[H\033[2J")
	return nil
}

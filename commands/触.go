package commands

import (
	"errors"
	"os"
	"time"
)

func init() {
	Register(&touchCmd{})
}

type touchCmd struct{}

func (t *touchCmd) Name() string        { return "触" }
func (t *touchCmd) Pinyin() string      { return "chù" }
func (t *touchCmd) Description() string { return "create an empty file or update its timestamp" }
func (t *touchCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("用法: 触 <file>  (usage: 触 <file>)")
	}
	for _, path := range args {
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		f.Close()
		now := time.Now()
		if err := os.Chtimes(path, now, now); err != nil {
			return err
		}
	}
	return nil
}

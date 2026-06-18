package commands

import (
	"fmt"
	"io/fs"
	"os"
	"time"
)

func init() {
	Register(&lsCmd{})
}

type lsCmd struct{}

func (l *lsCmd) Name() string        { return "列出" }
func (l *lsCmd) Pinyin() string      { return "lièchū" }
func (l *lsCmd) Description() string { return "list files in current directory" }
func (l *lsCmd) Execute(args []string) error {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// include . and ..
	for _, name := range []string{"..", "."} {
		info, err := os.Stat(dir + "/" + name)
		if err == nil {
			printEntry(info, name)
		}
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		printEntry(info, entry.Name())
	}
	return nil
}

func printEntry(info fs.FileInfo, name string) {
	mode := info.Mode()
	size := info.Size()
	mod := info.ModTime().Format(time.DateTime)
	fmt.Printf("%s  %8d  %s  %s\n", mode, size, mod, name)
}

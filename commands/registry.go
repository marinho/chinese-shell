package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mattn/go-runewidth"
	"zhell/score"
)

// Command is the interface every zhell command must implement.
type Command interface {
	// Name returns the Chinese command string (e.g. "出口").
	Name() string
	// Pinyin returns the romanised pronunciation (e.g. "chūkǒu").
	Pinyin() string
	// Linux returns the equivalent Linux/Unix command (e.g. "exit").
	Linux() string
	// Description is a short help string shown in the shell.
	Description() string
	// Execute runs the command with the given arguments.
	Execute(args []string) error
}

// registry holds all registered commands keyed by their Chinese name.
var registry = map[string]Command{}

var appVersion = "dev"
var appScore *score.Store

// SetVersion sets the application version shown in commands like 关于.
func SetVersion(v string) { appVersion = v }

// Version returns the current application version.
func Version() string { return appVersion }

// SetScore makes the score store available to commands like 分数.
func SetScore(s *score.Store) { appScore = s }

// Score returns the current score store (may be nil).
func Score() *score.Store { return appScore }

// Register adds a command to the global registry.
// Typically called from an init() function in each command file.
func Register(cmd Command) {
	registry[cmd.Name()] = cmd
}

// Lookup returns the command for the given name and whether it was found.
func Lookup(name string) (Command, bool) {
	cmd, ok := registry[name]
	return cmd, ok
}

// LookupByLinux returns the command whose Linux() value matches the given string.
func LookupByLinux(linux string) (Command, bool) {
	for _, cmd := range registry {
		if cmd.Linux() == linux {
			return cmd, true
		}
	}
	return nil, false
}

// All returns a copy of the full command registry.
func All() map[string]Command {
	out := make(map[string]Command, len(registry))
	for k, v := range registry {
		out[k] = v
	}
	return out
}

// PrintTable prints all registered commands as a formatted table.
func PrintTable() {
	all := All()
	keys := make([]string, 0, len(all))
	for k := range all {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	const colChinese, colPinyin, colLinux = 14, 22, 14
	pad := func(s string, width int) string {
		w := runewidth.StringWidth(s)
		if w >= width {
			return s
		}
		return s + strings.Repeat(" ", width-w)
	}

	fmt.Printf("%s  %s  %s  %s\n", pad("Chinese", colChinese), pad("Pinyin", colPinyin), pad("Linux", colLinux), "Description")
	fmt.Println(strings.Repeat("-", colChinese+colPinyin+colLinux+24))
	for _, k := range keys {
		cmd := all[k]
		fmt.Printf("%s  %s  %s  %s\n", pad(cmd.Name(), colChinese), pad(cmd.Pinyin(), colPinyin), pad(cmd.Linux(), colLinux), cmd.Description())
	}
}

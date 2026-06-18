package commands

// Command is the interface every zhell command must implement.
type Command interface {
	// Name returns the Chinese command string (e.g. "出口").
	Name() string
	// Pinyin returns the romanised pronunciation (e.g. "chūkǒu").
	Pinyin() string
	// Description is a short help string shown in the shell.
	Description() string
	// Execute runs the command with the given arguments.
	Execute(args []string) error
}

// registry holds all registered commands keyed by their Chinese name.
var registry = map[string]Command{}

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

// All returns a copy of the full command registry.
func All() map[string]Command {
	out := make(map[string]Command, len(registry))
	for k, v := range registry {
		out[k] = v
	}
	return out
}

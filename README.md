# zhell

A shell that accepts commands written in Chinese characters. This is just a playground project with no high expectations, but just to play with while I learn Chinese language.



## Usage

```
make
```

This builds and launches the REPL. At the `zhell>` prompt, type a Chinese command and press Enter.

```
欢迎使用 zhell！输入 '出口' 退出。
Welcome to zhell! Type '出口' to exit.

zhell> 出口
再见！感谢使用 zhell！
Goodbye! Thanks for using zhell!
```

## Commands

| Command | Equivalent in other shells | Meaning |
|---------|-----------------|-----------------|
| 出口    | exit | Exit the shell  |

## Adding a command

1. Create a new file under [commands/](commands/) (name it anything, e.g. `commands/帮助.go`).
2. Define a type that implements the `Command` interface (`Name`, `Description`, `Execute`).
3. Call `Register(&yourCmd{})` inside an `init()` function.

No changes to `main.go` are needed — the import of `zhell/commands` runs all `init()` functions automatically.

```go
package commands

import "fmt"

func init() { Register(&helpCmd{}) }

type helpCmd struct{}

func (h *helpCmd) Name() string        { return "帮助" }
func (h *helpCmd) Description() string { return "显示帮助 (show help)" }
func (h *helpCmd) Execute(_ []string) error {
    for _, cmd := range All() {
        fmt.Printf("  %s — %s\n", cmd.Name(), cmd.Description())
    }
    return nil
}
```

## Requirements

- Go 1.22+

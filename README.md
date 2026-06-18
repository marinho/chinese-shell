# zhell

A shell that accepts commands written in Chinese characters. This is just a playground project with no high expectations, but just to play with while I learn Chinese language.

## ⚠️ AI Warning ⚠️

This project has been mostly and shamelessly vibe-coded. Be careful and lower your expectations.

## Usage

```
make
```

This builds and launches the REPL. At the `zhell>` prompt, type a Chinese command and press Enter.

```
zhell> 这个目录
/Users/you/projects/zhell

zhell> 运行 /usr/bin/ls -la

zhell> 出口
再见！感谢使用 zhell！
Goodbye! Thanks for using zhell!
```

## Commands

| Chinese  | Pinyin       | Equivalent | Description                                    |
|----------|--------------|------------|------------------------------------------------|
| 帮       | bāng         | `help`     | Show available commands                        |
| 出口     | chūkǒu       | `exit`     | Exit the shell                                 |
| 关于     | guānyú       | —          | Show author and project information            |
| 日期     | rìqī         | `date`     | Print current date and time                    |
| 列出     | lièchū       | `ls -la`   | List files in current directory                |
| 这个目录 | zhège mùlù   | `pwd`      | Print working directory                        |
| 进入     | jìnrù        | `cd`       | Change current directory                       |
| 移动     | yídòng       | `mv`       | Move or rename a file or directory             |
| 复制     | fùzhì        | `cp`       | Copy a file or directory to another path       |
| 删除     | shānchú      | `rm`       | Remove a file or directory                     |
| 新目录   | xīn mùlù    | `mkdir`    | Create a new directory                         |
| 运行     | yùnxíng      | —          | Run a program at the given path                |
| 猫      | māo          | `cat`      | Print contents of a file                       |
| 触      | chù          | `touch`    | Create an empty file or update its timestamp   |
| 多      | duō          | `more`     | View file contents with a pager (forward only) |
| 少      | shǎo         | `less`     | View file contents with a pager                |
| 编辑     | biānjí       | `vim`      | Edit a file in vim                             |

## Adding a command

1. Create a new file under [commands/](commands/) (name it anything, e.g. `commands/帮助.go`).
2. Define a type that implements the `Command` interface (`Name`, `Pinyin`, `Description`, `Execute`).
3. Call `Register(&yourCmd{})` inside an `init()` function.

No changes to `main.go` are needed — the import of `zhell/commands` runs all `init()` functions automatically.

```go
package commands

import "fmt"

func init() { Register(&helpCmd{}) }

type helpCmd struct{}

func (h *helpCmd) Name() string        { return "帮助" }
func (h *helpCmd) Pinyin() string      { return "bāngzhù" }
func (h *helpCmd) Description() string { return "show help" }
func (h *helpCmd) Execute(_ []string) error {
    for _, cmd := range All() {
        fmt.Printf("  %-12s  %-20s  %s\n", cmd.Name(), cmd.Pinyin(), cmd.Description())
    }
    return nil
}
```

## Requirements

- Go 1.22+

# zhell

A shell that accepts commands written in Chinese characters. This is just a playground project with no high expectations, but just to play with while I learn Chinese language.

## ⚠️ AI Warning ⚠️

This project has been mostly and shamelessly vibe-coded. Be careful and lower your expectations.

## Usage

### Interactive shell

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

Type `帮` at the prompt to see all available commands.

### Inline command

Run a single command without entering the REPL:

```
./bin/zhell -c "日期"
./bin/zhell -c "列出 /tmp"
```

### Script mode

Save a script with a `.zh` extension and run it directly:

```
./bin/zhell my-script.zh
```

Scripts support `#` comments and shebang lines (`#!/usr/bin/env zhell`), so you can make them executable:

```bash
#!/usr/bin/env zhell
# hello-world.zh
说 你好，世界！
日期
```

### CLI options

| Option         | Description                        |
|----------------|------------------------------------|
| `-c <command>` | Execute a command string inline    |
| `--help`       | Show usage information             |
| `--帮`         | Show usage information (Chinese)   |
| `<file.zh>`    | Run a `.zh` script file            |

## English command suggestions

If you type a command in English that has a Chinese equivalent, zhell will suggest the Chinese command:

```
zhell> cat README.md
do you mean 猫 [māo]?
```

## Gamification

zhell tracks your score as you use it. The prompt shows your all-time score: `zhell [42]>`. Run `分数` at any time to see today/week/all-time totals and the full scoring rules.

| Event                                      | Points          |
|--------------------------------------------|-----------------|
| Correct Chinese command                    | +1              |
| First time using a command today           | +2 bonus        |
| Combo streak (consecutive different cmds)  | +N per step     |
| Typing an English equivalent               | −1, reset combo |
| Unknown command                            | −1, reset combo |

- Repeating the same command back-to-back scores no points and resets your combo streak.
- Calling `帮` or `分数` never scores or deducts points.
- Score never drops below 0.
- Today's score, this week's score, and all-time total are all tracked separately in `~/.zhell_score.json`.

## Commands

| Chinese  | Pinyin       | Linux           | Description                                    |
|----------|--------------|-----------------|------------------------------------------------|
| 帮       | bāng         | `help`          | Show available commands                        |
| 出口     | chūkǒu       | `exit`          | Exit the shell                                 |
| 关于     | guānyú       | —               | Show author and project information            |
| 清屏     | qīng píng    | `clear`         | Clear the screen                               |
| 日期     | rìqī         | `date`          | Print current date and time                    |
| 列出     | lièchū       | `ls`            | List files in current directory                |
| 这个目录 | zhège mùlù   | `pwd`           | Print working directory                        |
| 进入     | jìnrù        | `cd`            | Change current directory                       |
| 移动     | yídòng       | `mv`            | Move or rename a file or directory             |
| 复制     | fùzhì        | `cp`            | Copy a file or directory to another path       |
| 删除     | shānchú      | `rm`            | Remove a file or directory                     |
| 新目录   | xīn mùlù     | `mkdir`         | Create a new directory                         |
| 运行     | yùnxíng      | `./`            | Run a program at the given path                |
| 猫       | māo          | `cat`           | Print contents of a file                       |
| 新文件   | xīn wénjiàn  | `touch`         | Create a new empty file                        |
| 多       | duō          | `more`          | View file contents with a pager (forward only) |
| 少       | shǎo         | `less`          | View file contents with a pager                |
| 编辑     | biānjí       | `vim`           | Edit a file in vim                             |
| 蠢货     | chǔnhuò      | `git`           | Run git with the given arguments               |
| 改权限   | gǎi quánxiàn | `chmod`         | Change file permissions                        |
| 改群组   | gǎi qúnzǔ    | `chgrp`         | Change file group                              |
| 改归属   | gǎi guīshǔ   | `chown`         | Change file owner                              |
| 看归属   | kàn guīshǔ   | `stat`          | Show owner and group of a file                 |
| 列用户   | liè yònghù   | `getent passwd` | List all users on the system                   |
| 列群组   | liè qúnzǔ    | `getent group`  | List all groups on the system                  |
| 我是谁   | wǒ shì shuí  | `whoami`        | Print current user                             |
| 分数     | fēnshù       | `score`         | Show your score and how the scoring system works |
| 提权     | tí quán      | `sudo`          | Run a zhell command as superuser               |
| 设       | shè          | `export`        | Set an environment variable                    |
| 说       | shuō         | `echo`          | Print arguments to stdout                      |

## Environment variables

Set a variable with `设` and reference it anywhere with `$VAR`:

```
zhell> 设 NAME 世界
zhell> 说 你好，$NAME！
你好，世界！

zhell> 进入 $HOME
```

Variable expansion happens before the command is parsed, so `$VAR` works as an argument to any command. Standard system variables like `$HOME` and `$PATH` are also available.

## Adding a command

1. Create a new file under [commands/](commands/) (name it anything, e.g. `commands/帮助.go`).
2. Define a type that implements the `Command` interface (`Name`, `Pinyin`, `Linux`, `Description`, `Execute`).
3. Call `Register(&yourCmd{})` inside an `init()` function.

No changes to `main.go` are needed — the import of `zhell/commands` runs all `init()` functions automatically.

```go
package commands

import "fmt"

func init() { Register(&helpCmd{}) }

type helpCmd struct{}

func (h *helpCmd) Name() string        { return "帮助" }
func (h *helpCmd) Pinyin() string      { return "bāngzhù" }
func (h *helpCmd) Linux() string       { return "help" }
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

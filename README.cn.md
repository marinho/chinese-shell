# zhell

一个接受中文字符命令的 Shell。这只是一个没有太高期望的练习项目，用来在学习中文的同时玩一玩。

## ⚠️ AI 警告 ⚠️

本项目大部分内容由 AI 辅助生成，毫无羞耻之心。请谨慎使用，降低预期。

## 使用方法

```
make
```

这会编译并启动 REPL。在 `zhell>` 提示符下，输入中文命令并按回车。

```
欢迎使用 zhell！输入 '出口' 退出。
Welcome to zhell! Type '出口' to exit.

Chinese       Pinyin                Description
--------------------------------------------------
帮            bāng                  show available commands
出口          chūkǒu                exit the shell
日期          rìqī                  print current date and time
列出          lièchū                list files in current directory
进入          jìnrù                 change current directory
猫            māo                   print contents of a file
这个目录       zhège mùlù            print working directory
运行          yùnxíng               run a program at the given path

zhell> 这个目录
/Users/you/projects/zhell

zhell> 运行 /usr/bin/ls -la

zhell> 出口
再见！感谢使用 zhell！
Goodbye! Thanks for using zhell!
```

## 命令

| 中文     | 拼音         | 等价命令   | 说明                        |
|----------|--------------|------------|-----------------------------|
| 帮       | bāng         | `help`     | 显示可用命令                |
| 出口     | chūkǒu       | `exit`     | 退出 Shell                  |
| 日期     | rìqī         | `date`     | 打印当前日期和时间          |
| 列出     | lièchū       | `ls -la`   | 列出当前目录的文件          |
| 进入     | jìnrù        | `cd`       | 切换当前目录                |
| 猫      | māo          | `cat`      | 打印文件内容                |
| 这个目录 | zhège mùlù   | `pwd`      | 打印当前目录                |
| 运行     | yùnxíng      | —          | 运行指定路径的程序          |

## 添加命令

1. 在 [commands/](commands/) 目录下新建一个文件（文件名随意，例如 `commands/帮助.go`）。
2. 定义一个实现 `Command` 接口的类型（需实现 `Name`、`Pinyin`、`Description`、`Execute`）。
3. 在 `init()` 函数中调用 `Register(&yourCmd{})`。

无需修改 `main.go`——导入 `zhell/commands` 包时会自动执行所有 `init()` 函数完成注册。

```go
package commands

import "fmt"

func init() { Register(&helpCmd{}) }

type helpCmd struct{}

func (h *helpCmd) Name() string        { return "帮助" }
func (h *helpCmd) Pinyin() string      { return "bāngzhù" }
func (h *helpCmd) Description() string { return "显示帮助" }
func (h *helpCmd) Execute(_ []string) error {
    for _, cmd := range All() {
        fmt.Printf("  %-12s  %-20s  %s\n", cmd.Name(), cmd.Pinyin(), cmd.Description())
    }
    return nil
}
```

## 环境要求

- Go 1.22+

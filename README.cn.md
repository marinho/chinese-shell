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
zhell> 这个目录
/Users/you/projects/zhell

zhell> 运行 /usr/bin/ls -la

zhell> 出口
再见！感谢使用 zhell！
Goodbye! Thanks for using zhell!
```

## 命令

| 中文     | 拼音         | 等价命令   | 说明                              |
|----------|--------------|------------|-----------------------------------|
| 帮       | bāng         | `help`     | 显示可用命令                      |
| 出口     | chūkǒu       | `exit`     | 退出 Shell                        |
| 关于     | guānyú       | —          | 显示作者和项目信息                |
| 清屏     | qīng píng    | `clear`    | 清除屏幕                          |
| 日期     | rìqī         | `date`     | 打印当前日期和时间                |
| 列出     | lièchū       | `ls -la`   | 列出当前目录的文件                |
| 进入     | jìnrù        | `cd`       | 切换当前目录                      |
| 猫      | māo          | `cat`      | 打印文件内容                      |
| 新文件   | xīn wénjiàn  | `touch`    | 创建新空文件                      |
| 少      | shǎo         | `less`     | 用分页器查看文件内容              |
| 多      | duō          | `more`     | 用分页器查看文件内容（只能向前）  |
| 编辑     | biānjí       | `vim`      | 用 vim 编辑文件                   |
| 蠢货     | chǔnhuò      | `git`      | 运行 git 命令                     |
| 改权限   | gǎi quánxiàn | `chmod`    | 修改文件权限                      |
| 改群组   | gǎi qúnzǔ    | `chgrp`    | 修改文件群组                      |
| 改归属   | gǎi guīshǔ   | `chown`    | 修改文件归属者                    |
| 列用户   | liè yònghù   | —          | 列出系统所有用户                  |
| 列群组   | liè qúnzǔ    | —          | 列出系统所有群组                  |
| 我是谁   | wǒ shì shuí  | `whoami`   | 打印当前用户                      |
| 移动     | yídòng       | `mv`       | 移动或重命名文件或目录            |
| 复制     | fùzhì        | `cp`       | 复制文件或目录到另一路径          |
| 删除     | shānchú      | `rm`       | 删除文件或目录                    |
| 新目录   | xīn mùlù    | `mkdir`    | 创建新目录                        |
| 这个目录 | zhège mùlù   | `pwd`      | 打印当前目录                      |
| 运行     | yùnxíng      | —          | 运行指定路径的程序                |

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

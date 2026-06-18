package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/chzyer/readline"

	// Importing the commands package causes all init() functions in
	// that package to run, registering every command automatically.
	// Add a new file under commands/ and call Register() in its init()
	// — no changes to main.go are needed.
	"zhell/commands"
)

const prompt = "zhell> "

func main() {
	commands.SetVersion(version)
	fmt.Printf("欢迎使用 zhell v%s！输入 '出口' 退出。\n", version)
	fmt.Printf("Welcome to zhell v%s! Type '出口' to exit.\n", version)
	fmt.Println()
	commands.PrintTable()
	fmt.Println()

	historyFile := filepath.Join(os.Getenv("HOME"), ".zhell_history")

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          prompt,
		HistoryFile:     historyFile,
		HistoryLimit:    100,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "zhell: readline init error: %v\n", err)
		os.Exit(1)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			continue
		}
		if err == io.EOF {
			fmt.Println()
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "zhell: %v\n", err)
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		name, args := parts[0], parts[1:]

		cmd, ok := commands.Lookup(name)
		if !ok {
			fmt.Fprintf(os.Stderr, "zhell: 未知命令: %s\n", name)
			continue
		}

		if err := cmd.Execute(args); err != nil {
			fmt.Fprintf(os.Stderr, "zhell: 错误: %v\n", err)
		}
	}
}

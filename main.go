package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// Importing the commands package causes all init() functions in
	// that package to run, registering every command automatically.
	// Add a new file under commands/ and call Register() in its init()
	// — no changes to main.go are needed.
	"zhell/commands"
)

const prompt = "zhell> "

func main() {
	fmt.Println("欢迎使用 zhell！输入 '出口' 退出。")
	fmt.Println("Welcome to zhell! Type '出口' to exit.")
	fmt.Println()
	commands.PrintTable()
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			// EOF (Ctrl-D)
			fmt.Println()
			break
		}

		line := strings.TrimSpace(scanner.Text())
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

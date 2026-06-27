package main

import (
	"bufio"
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

func executeLine(line string) error {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") {
		return nil
	}
	parts := strings.Fields(line)
	name, args := parts[0], parts[1:]
	cmd, ok := commands.Lookup(name)
	if !ok {
		return fmt.Errorf("未知命令: %s", name)
	}
	return cmd.Execute(args)
}

func runScript(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "zhell: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.HasPrefix(line, "#!") {
			continue // skip shebang
		}
		if err := executeLine(line); err != nil {
			fmt.Fprintf(os.Stderr, "zhell: %s:%d: %v\n", path, lineNum, err)
			os.Exit(1)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "zhell: 读取脚本错误: %v\n", err)
		os.Exit(1)
	}
}

// pathCompleter completes filesystem paths for the last argument on the line.
type pathCompleter struct{}

func (p *pathCompleter) Do(line []rune, pos int) ([][]rune, int) {
	// work only on the segment up to the cursor
	input := string(line[:pos])
	parts := strings.Fields(input)

	var prefix string
	if len(parts) == 0 || (len(parts) == 1 && !strings.HasSuffix(input, " ")) {
		// still typing the command name — no path completion
		return nil, 0
	}
	if strings.HasSuffix(input, " ") {
		prefix = ""
	} else {
		prefix = parts[len(parts)-1]
	}

	dir, file := filepath.Split(prefix)
	searchDir := dir
	if searchDir == "" {
		searchDir = "."
	}

	entries, err := os.ReadDir(searchDir)
	if err != nil {
		return nil, 0
	}

	var matches []string
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, file) {
			full := dir + name
			if e.IsDir() {
				full += "/"
			}
			matches = append(matches, full)
		}
	}

	if len(matches) == 0 {
		return nil, 0
	}

	// strip the already-typed prefix so readline inserts only the completion
	completions := make([][]rune, len(matches))
	for i, m := range matches {
		completions[i] = []rune(m[len(prefix):])
	}
	return completions, len([]rune(prefix))
}

func main() {
	commands.SetVersion(version)

	if len(os.Args) > 1 {
		path := os.Args[1]
		if ext := filepath.Ext(path); ext != ".zh" {
			fmt.Fprintf(os.Stderr, "zhell: 脚本文件必须以 .zh 结尾 (script file must have .zh extension)\n")
			os.Exit(1)
		}
		runScript(path)
		return
	}

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
		AutoComplete:    &pathCompleter{},
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

		if err := executeLine(line); err != nil {
			fmt.Fprintf(os.Stderr, "zhell: 错误: %v\n", err)
		}
	}
}

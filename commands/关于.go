package commands

import "fmt"

func init() {
	Register(&aboutCmd{})
}

type aboutCmd struct{}

func (a *aboutCmd) Name() string        { return "关于" }
func (a *aboutCmd) Pinyin() string      { return "guānyú" }
func (a *aboutCmd) Linux() string       { return "-" }
func (a *aboutCmd) Description() string { return "show author and project information" }
func (a *aboutCmd) Execute(_ []string) error {
	fmt.Printf("zhell v%s — a shell that accepts commands written in Chinese characters\n", Version())
	fmt.Println()
	fmt.Println("Author:  Marinho Brandao")
	fmt.Println("GitHub:  https://github.com/marinho/chinese-shell/")
	return nil
}

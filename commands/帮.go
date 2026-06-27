package commands

func init() {
	Register(&helpCmd{})
}

type helpCmd struct{}

func (h *helpCmd) Name() string        { return "帮" }
func (h *helpCmd) Pinyin() string      { return "bāng" }
func (h *helpCmd) Linux() string       { return "help" }
func (h *helpCmd) Description() string { return "show available commands" }
func (h *helpCmd) Execute(_ []string) error {
	PrintTable()
	return nil
}

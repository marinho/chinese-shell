package commands

import (
	"fmt"
	"os/user"
	"strings"
)

func init() {
	Register(&listUsersCmd{})
}

type listUsersCmd struct{}

func (c *listUsersCmd) Name() string        { return "列用户" }
func (c *listUsersCmd) Pinyin() string      { return "liè yònghù" }
func (c *listUsersCmd) Description() string { return "list all users on the system" }
func (c *listUsersCmd) Execute(args []string) error {
	users, err := getAllUsers()
	if err != nil {
		return err
	}
	fmt.Printf("%-20s  %-6s  %s\n", "用户名", "UID", "全名")
	for _, u := range users {
		fmt.Printf("%-20s  %-6s  %s\n", u.Username, u.Uid, u.Name)
	}
	return nil
}

func getAllUsers() ([]*user.User, error) {
	// macOS stores users in Directory Services, not /etc/passwd
	out, err := runCommand("dscl", ".", "-list", "/Users")
	if err == nil {
		var users []*user.User
		for _, name := range strings.Split(strings.TrimSpace(out), "\n") {
			if name == "" || strings.HasPrefix(name, "_") {
				continue
			}
			u, err := user.Lookup(name)
			if err != nil {
				continue
			}
			users = append(users, u)
		}
		return users, nil
	}
	// fallback for Linux
	data, err := readLines("/etc/passwd")
	if err != nil {
		return nil, err
	}
	var users []*user.User
	for _, line := range data {
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		fields := splitColon(line, 7)
		if len(fields) < 4 {
			continue
		}
		u, err := user.LookupId(fields[2])
		if err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

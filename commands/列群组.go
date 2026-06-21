package commands

import (
	"fmt"
	"os/user"
)

func init() {
	Register(&listGroupsCmd{})
}

type listGroupsCmd struct{}

func (c *listGroupsCmd) Name() string        { return "列群组" }
func (c *listGroupsCmd) Pinyin() string      { return "liè qúnzǔ" }
func (c *listGroupsCmd) Description() string { return "list all groups on the system" }
func (c *listGroupsCmd) Execute(args []string) error {
	groups, err := getAllGroups()
	if err != nil {
		return err
	}
	fmt.Printf("%-20s  %s\n", "群组名", "GID")
	for _, g := range groups {
		fmt.Printf("%-20s  %s\n", g.Name, g.Gid)
	}
	return nil
}

func getAllGroups() ([]*user.Group, error) {
	data, err := readLines("/etc/group")
	if err != nil {
		return nil, err
	}
	var groups []*user.Group
	for _, line := range data {
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		// fields: name:password:gid:members
		fields := splitColon(line, 4)
		if len(fields) < 3 {
			continue
		}
		g, err := user.LookupGroupId(fields[2])
		if err != nil {
			continue
		}
		groups = append(groups, g)
	}
	return groups, nil
}

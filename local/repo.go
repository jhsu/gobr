package gobr

import (
	"bytes"
	"os/exec"
	"strings"
)

// Get a range of local branches.
func Branches() []string {
	cmd := exec.Command("git", "branch")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	branches := strings.Fields(out.String())
	names := []string{}
	for _, name := range branches {
		if name != "" && name != "*" {
			names = append(names, name)
		}
	}
	return names
}

func SetBranch(branch string) {
	cmd := exec.Command("git", "checkout", branch)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

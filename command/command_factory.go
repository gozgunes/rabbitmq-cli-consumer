package command

import (
	"os/exec"
)

type CommandFactory struct {
	Cmd  string
	Args []string
}

func (me CommandFactory) Create(body string) *exec.Cmd {
	cmd := exec.Command(me.Cmd, me.Args...)

	stdin := cmd.StdinPipe()

	go func() {
		if(len(body) > 0) {
			defer stdin.Close()
			io.WriteString(stdin, body)
		}
	}()

	return cmd
}

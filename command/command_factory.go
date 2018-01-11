package command

import (
	"os/exec"
		"log"
    	"io"
)

type CommandFactory struct {
	Cmd  string
	Args []string
}

func (me CommandFactory) Create(body string) *exec.Cmd {
	cmd := exec.Command(me.Cmd, me.Args...)

	stdin, err := cmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if(len(body) > 0) {
			defer stdin.Close()
			io.WriteString(stdin, body)
		}
	}()

	return cmd
}

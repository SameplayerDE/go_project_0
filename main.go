package main

import (
	"bufio"
	"os"
	"strings"
)

type CommandExecutor interface {
	execute() (string, int)
}

type ConsoleCommand struct {
	command string
	args    []string
	hasArgs bool
}

func readCommand(reader *bufio.Reader) []string {
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")
	command := strings.Split(text, " ")
	finalCommand := make([]string, 0)
	for i := 0; i < len(command); i++ {
		a := command[i]
		if len(a) <= 0 {
			continue
		}
		a = strings.Replace(a, " ", "", -1)
		finalCommand = append(finalCommand, a)
	}
	return finalCommand
}

func (cmd *ConsoleCommand) Parse(args []string) {
	cmd.command = args[0]
	if len(args) > 1 {
		cmd.args = args[1:]
	} else {
		cmd.args = nil
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	running := true
	var cmd ConsoleCommand
	for running {
		args := readCommand(reader)
		if len(args) > 0 {
			cmd.Parse(args)
			if cmd.command == "!" {
				if cmd.args == nil {
					running = false
				}
			}
		}
	}
}

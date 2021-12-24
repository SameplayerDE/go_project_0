package main

import (
	"bufio"
	"os"
	"strings"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	running := true
	for running {
		args := readCommand(reader)
		if len(args) > 0 {
			if len(args) == 1 {
				if args[0] == "!" {
					running = false
				}
			}
		}
	}
}

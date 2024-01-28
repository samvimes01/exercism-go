package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Command struct {
	Name string
	Mp   int
	W    int
	D    int
	L    int
	P    int
}

func Tally(reader io.Reader, writer io.Writer) error {
	input := []string{}
	r := bufio.NewReader(reader)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if len(line) != 0 && line[0] != '#' {
			input = append(input, string(line))
		}
	}
	commands, err := getCommands(input)
    if err != nil {
        return err
    }

	commandsOrder := sortCommands(commands)

	output := []string{fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s", "Team", "MP", "W", "D", "L", "P")}
	tpl := "%-30s | %2d | %2d | %2d | %2d | %2d"
	for _, c := range commandsOrder {
		output = append(output, fmt.Sprintf(tpl, c.Name, c.Mp, c.W, c.D, c.L, c.P))
	}
	fmt.Fprintln(writer, strings.Join(output, "\n"))
	return nil
}

func sortCommands(commands map[string]*Command) []*Command {
	commandsOrder := make([]*Command, 0, len(commands))
	for _, c := range commands {
		commandsOrder = append(commandsOrder, c)
	}
	sort.Slice(commandsOrder, func(i, j int) bool {
		if commandsOrder[i].P == commandsOrder[j].P {
			return commandsOrder[i].Name < commandsOrder[j].Name
		}
		return commandsOrder[i].P > commandsOrder[j].P
	})
	return commandsOrder
}

func getCommands(input []string) (map[string]*Command, error) {
	commands := make(map[string]*Command)
	for _, line := range input {
		res := strings.Split(line, ";")
        if len(res) != 3 {
			return nil, fmt.Errorf("invalid input")
		}
		c1, c2, r := strings.TrimSpace(res[0]), strings.TrimSpace(res[1]), strings.TrimSpace(res[2])
        if c1 == c2 {
			return nil, fmt.Errorf("same team")
		}
        if r != "win" && r != "loss" && r != "draw" {
            return nil, fmt.Errorf("unknown result")
		}
		if _, ok := commands[c1]; !ok {
			commands[c1] = &Command{Name: c1}
		}
		if _, ok := commands[c2]; !ok {
			commands[c2] = &Command{Name: c2}
		}
		commands[c1].Mp++
		commands[c2].Mp++
		switch r {
		case "win":
			commands[c1].W++
			commands[c1].P += 3
			commands[c2].L++
		case "loss":
			commands[c2].W++
			commands[c2].P += 3
			commands[c1].L++
		case "draw":
			commands[c1].D++
			commands[c1].P++
			commands[c2].D++
			commands[c2].P++
		}
	}
	return commands, nil
}

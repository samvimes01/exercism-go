package kindergarten

import (
	"errors"
	"slices"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

var plantsDict = map[rune]string{
	'V': "violets",
	'G': "grass",
	'C': "clover",
	'R': "radishes",
}

type Garden struct {
	names  []string
	rows   [][]string
	plants []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	lines := strings.Split(diagram, "\n")
	var rowCods []string
	for _, line := range lines {
		if line != "" {
			rowCods = append(rowCods, line)
		}
	}
	if len(children) == 0 || len(diagram) == 0 || len(lines) == len(rowCods) || len(rowCods) < 2 || len(rowCods[0]) != len(rowCods[1]) || (len(rowCods[0])%2 != 0 || len(rowCods[1])%2 != 0) {
		return nil, errors.New("invalid garden")
	}
	
	childMap := make(map[string]struct{})
	for _, child := range children {
		if _, ok := childMap[child]; ok {
			return nil, errors.New("duplicate child")
		}
		childMap[child] = struct{}{}
	}

	names := make([]string, len(children))
	copy(names, children)
	sort.Strings(names)

	rows := make([][]string, len(rowCods))
	plants := make([]string, 0)
	for i, row := range rowCods {
		rows[i] = make([]string, 0)
		for _, p := range row {
			if plant, ok := plantsDict[p]; ok {
				rows[i] = append(rows[i], plant)
				plants = append(plants, plant)
			} else {
				return nil, errors.New("invalid plant")
			}
		}
	}

	return &Garden{names, rows, plants}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	i := slices.Index(g.names, child)
	if i == -1 {
		return g.plants, false
	}
	plants := make([]string, 0)
	for _, r := range g.rows {
		plants = append(plants, r[i*2:i*2+2]...)
	}
	return plants, true
}

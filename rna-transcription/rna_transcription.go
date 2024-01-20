package strand

import "strings"

func ToRNA(dna string) string {
	dnaToRnaMap := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	rna := strings.Builder{}
	for _, n := range dna {
		rna.WriteRune(dnaToRnaMap[n])
	}
	return rna.String()
}

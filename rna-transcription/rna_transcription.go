// Package strand gives methods to work with DNA Strands.
package strand

import "bytes"

var rnaTable = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA takes an DNA and return it's match to the RNA table.
func ToRNA(dna string) string {
	buf := bytes.Buffer{}
	for _, rune := range dna {
		buf.WriteString(rnaTable[string(rune)])
	}
	return buf.String()
}

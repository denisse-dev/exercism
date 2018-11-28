// Package strand returns the RNA complement of a DNA strand
package strand

// ToRNA changes the four nucleotids found in DNA to those found in RNA
func ToRNA(dna string) (rna string) {
	for i := range dna {
		switch dna[i] {
		case 'G':
			rna = rna + "C"
		case 'C':
			rna = rna + "G"
		case 'T':
			rna = rna + "A"
		case 'A':
			rna = rna + "U"
		}
	}
	return
}

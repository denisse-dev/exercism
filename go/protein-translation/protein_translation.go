package protein

import "errors"

var (
	// ErrStop is returned when a STOP codon is found
	ErrStop = errors.New("STOP")
	// ErrInvalidBase is returned when an invalid codon is found
	ErrInvalidBase = errors.New("Invalid")
)

// FromCodon matches a codon to a protein
func FromCodon(rna string) (string, error) {
	switch rna {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA creates a slice codons from an RNA sequence
func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna); i += 3 {
		rnaSegment := rna[i : i+3]
		if protein, err := FromCodon(rnaSegment); err == nil {
			proteins = append(proteins, protein)
		} else if err == ErrStop {
			return proteins, nil
		} else {
			return proteins, err
		}
	}
	return proteins, nil
}

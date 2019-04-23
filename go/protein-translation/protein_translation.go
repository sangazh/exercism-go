package protein

func FromCodon(codon string) string {
	CodonProteinMap := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	return CodonProteinMap[codon]

}

func FromRNA(rna string) []string {
	var result []string
	for i := 0; i < len(rna); i = i + 3 {
		codon := rna[i : i+3]
		protein := FromCodon(codon)
		if protein == "STOP" {
			break
		}
		result = append(result, protein)
	}

	return result
}

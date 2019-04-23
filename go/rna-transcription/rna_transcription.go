package strand

func ToRNA(dna string) (rna string) {
	transcript := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}

	for _, n := range dna {
		if v, ok := transcript[string(n)]; ok {
			rna += v
		}
	}

	return
}

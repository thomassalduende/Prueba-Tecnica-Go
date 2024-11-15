package service

func IsMutant(dna []string) bool {
	N := len(dna)
	if N == 0 {
		return false
	}
	var count int

	// Verificando horizontalmente y verticalmente
	for i := 0; i < N; i++ {
		for j := 0; j <= N-4; j++ {
			if dna[i][j] == dna[i][j+1] && dna[i][j] == dna[i][j+2] && dna[i][j] == dna[i][j+3] {
				count++
				if count > 1 {
					return true
				}
			}
			if dna[j][i] == dna[j+1][i] && dna[j][i] == dna[j+2][i] && dna[j][i] == dna[j+3][i] {
				count++
				if count > 1 {
					return true
				}
			}
		}
	}

	// Verificando diagonalmente
	for i := 0; i <= N-4; i++ {
		for j := 0; j <= N-4; j++ {
			if dna[i][j] == dna[i+1][j+1] && dna[i][j] == dna[i+2][j+2] && dna[i][j] == dna[i+3][j+3] {
				count++
				if count > 1 {
					return true
				}
			}
			if dna[i][N-1-j] == dna[i+1][N-2-j] && dna[i][N-1-j] == dna[i+2][N-3-j] && dna[i][N-1-j] == dna[i+3][N-4-j] {
				count++
				if count > 1 {
					return true
				}
			}
		}
	}

	return false
}

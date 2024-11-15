package model

import (
	"context"
	"mutant-checker/database"
)

type DNA struct {
	Sequence string
	IsMutant bool
}

func InsertDNA(sequence string, isMutant bool) error {
	_, err := database.Conn.Exec(context.Background(),
		"INSERT INTO dna_sequences (sequence, is_mutant) VALUES ($1, $2) ON CONFLICT (sequence) DO NOTHING",
		sequence, isMutant)
	if err != nil {
		return err
	}
	return nil
}

func CountMutantDNA() (int, error) {
	var count int
	err := database.Conn.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM dna_sequences WHERE is_mutant = true").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountHumanDNA() (int, error) {
	var count int
	err := database.Conn.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM dna_sequences WHERE is_mutant = false").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

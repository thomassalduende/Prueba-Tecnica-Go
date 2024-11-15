package controller

import (
	"encoding/json"
	"fmt"
	"mutant-checker/models"
	"mutant-checker/services"
	"net/http"
)

type DnaRequest struct {
	Dna []string `json:"dna"`
}

type StatsResponse struct {
	CountMutantDna int     `json:"count_mutant_dna"`
	CountHumanDna  int     `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

func MutantHandler(w http.ResponseWriter, r *http.Request) {
	var req DnaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || len(req.Dna) == 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	isMutant := service.IsMutant(req.Dna)

	// Convertir la secuencia de ADN a un solo string para almacenar en la base de datos
	dnaSequence := fmt.Sprintf("%v", req.Dna)

	// Insertar la secuencia de ADN en la base de datos
	err = model.InsertDNA(dnaSequence, isMutant)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if isMutant {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Es un mutante"))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("No es un mutante"))
	}
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	countMutantDna, err := model.CountMutantDNA()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	countHumanDna, err := model.CountHumanDNA()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	ratio := float64(countMutantDna) / float64(countHumanDna)

	stats := StatsResponse{
		CountMutantDna: countMutantDna,
		CountHumanDna:  countHumanDna,
		Ratio:          ratio,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Conn *pgxpool.Pool

func InitDB() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	Conn, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Connected to the database!")

	// Crear la tabla si no existe
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS dna_sequences (
        id SERIAL PRIMARY KEY,
        sequence TEXT UNIQUE NOT NULL,
        is_mutant BOOLEAN NOT NULL
    );
    `

	_, err = Conn.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v\n", err)
	}
	fmt.Println("Table dna_sequences is ready!")
}

func CloseDB() {
	if Conn != nil {
		Conn.Close()
	}
}

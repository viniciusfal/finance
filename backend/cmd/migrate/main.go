package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbURL = "postgresql://postgres:fuTjLCHygdFJTpRARveKdgGtkwFOzpgc@mainline.proxy.rlwy.net:18337/railway"

func main() {
	// Conectar ao banco
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	// Ler arquivo de migration
	migrationPath := filepath.Join("..", "..", "migrations", "001_create_tables.sql")
	if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
		// Tentar caminho alternativo
		migrationPath = filepath.Join("migrations", "001_create_tables.sql")
	}

	sql, err := ioutil.ReadFile(migrationPath)
	if err != nil {
		log.Fatal("Failed to read migration file:", err)
	}

	// Executar migration
	ctx := context.Background()
	_, err = pool.Exec(ctx, string(sql))
	if err != nil {
		log.Fatal("Failed to execute migration:", err)
	}

	fmt.Println("Migration executed successfully!")
}

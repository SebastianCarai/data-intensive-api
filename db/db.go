package db

import (
	"context"
	"fmt"
	"log"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	dotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Pool *pgxpool.Pool

func ConnectDB(){
	envFile, envFileErr := dotenv.Read(".env.dev")
	if envFileErr != nil{
		log.Fatal(envFileErr)
	}

	pool, err := pgxpool.New(context.Background(), envFile["POSTGRES_URL"])
	
	if err != nil{
		fmt.Printf("Error while connecting to DB: %v\n", err)
	}

	Pool = pool
}
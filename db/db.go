package db

import (
	"database/sql"
	"fmt"
	"log"

	dotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func ConnectDB() *sql.DB{
	envFile, envFileErr := dotenv.Read(".env.dev")
	if envFileErr != nil{
		log.Fatal(envFileErr)
	}

	host := envFile["HOST"]
	dbPort := envFile["POSTGRES_PORT"]
	dbName:= envFile["POSTGRES_DB"]
	dbUser := envFile["POSTGRES_USER"]
	dbPassword := envFile["POSTGRES_PASSWORD"]

	psqlSetup := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, dbPort, dbUser, dbName, dbPassword)

	db, errSql := sql.Open("postgres", psqlSetup)

	if errSql != nil{
		log.Fatal(errSql)
	}

	return db
}
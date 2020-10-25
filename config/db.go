package db

import (
	"KlikDigital/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var GlobalDbSQL *sql.DB

func InitDB() (*sql.DB, error)  {
	dataConfig, err := utils.GetConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	connInfo := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=disable`, dataConfig.Database.Username, dataConfig.Database.Password, dataConfig.Database.Host, dataConfig.Database.Port, dataConfig.Database.Name)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatalln(err)
		return &sql.DB{}, nil
	}

	return db, nil
}

func init()  {
	GlobalDbSQL, _ = InitDB()
}
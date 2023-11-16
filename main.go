package main

import (
	"database/sql"
	"log"

	"github.com/gafar-code/quran-server/api"
	db "github.com/gafar-code/quran-server/db/sqlc"
	"github.com/gafar-code/quran-server/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Tidak dapat memuat config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DbSource)
	if err != nil {
		log.Fatal("Tidak bisa terhubung ke database:", err)
	}

	page := db.NewPage(conn)
	server := api.NewServer(page)

	server.Start(config.ServerAddress)
}

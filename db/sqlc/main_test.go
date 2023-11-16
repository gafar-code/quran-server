package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gafar-code/quran-server/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Tidak dapat memuat config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DbSource)
	if err != nil {
		log.Fatal("Tidak bisa terhubung ke database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

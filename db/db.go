package db
import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	dp, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return dp, nil
}
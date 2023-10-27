package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pooulad/csv-convertor/config"
	"github.com/pooulad/csv-convertor/internal/readflag"
)

func ConnectMysql(cfg *config.MysqlConfig, f *readflag.FlagReturns) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	fmt.Println("mysql connected")
	return db, nil
}


package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/pooulad/csv-convertor/config"
)

func ConnectMysql(cfg *config.MysqlConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
fmt.Println(dsn)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}

func ConnectPostgres(cfg *config.PostgresConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Name, cfg.SSL)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("postgres connected")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}

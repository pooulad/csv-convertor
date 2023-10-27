package main

import (
	"fmt"
	"log"

	"github.com/pooulad/csv-convertor/config"
	"github.com/pooulad/csv-convertor/database"
	"github.com/pooulad/csv-convertor/internal/readflag"
	"github.com/pooulad/csv-convertor/utils"
)

func main() {
	flags, err := readflag.ReadFlag()
	if err != nil {
		utils.Colorize(utils.ColorRed, err.Error())
		return
	}
	// var dbConfig interface{}
	if flags.DbType == "postgres" {
		config, err := config.ReadPostgresConfig(flags.DbConfig)
		if err != nil {
			fmt.Println(err)
			return
		}
		// dbConfig = config

		db, err := database.ConnectPostgres(config, flags)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

	} else if flags.DbType == "mysql"{
		config, err := config.ReadMysqlConfig(flags.DbConfig)
		if err != nil {
			fmt.Println(err)
			return
		}
		// dbConfig = config
		fmt.Println(config)
		db, err := database.ConnectMysql(config, flags)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}else{
		utils.Colorize(utils.ColorRed, "database name should be `postgres` OR `mysql`")
		return
	}


	fmt.Printf("this is config : %s\n",dbConfig)
}

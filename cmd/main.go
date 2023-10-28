package main

import (
	"log"

	"github.com/pooulad/csv-convertor/config"
	"github.com/pooulad/csv-convertor/database"
	"github.com/pooulad/csv-convertor/internal/readcsv"
	"github.com/pooulad/csv-convertor/internal/readflag"
	"github.com/pooulad/csv-convertor/utils"
)

func main() {
	flags, err := readflag.ReadFlag()
	if err != nil {
		utils.Colorize(utils.ColorRed, err.Error())
		return
	}
	if flags.DbType == "postgres" {
		config, err := config.ReadPostgresConfig(flags.DbConfig)

		if err != nil {
			utils.Colorize(utils.ColorRed, err.Error())
			return
		}

		db, err := database.ConnectPostgres(config)
		if err != nil {
			log.Fatal(err)
		}


		err = readcsv.ReadExcelAndInsertData(db, flags.FileName,config.Table)
		if err != nil {
			utils.Colorize(utils.ColorRed, err.Error())
			return
		}

		defer db.Close()

	} else if flags.DbType == "mysql" {
		config, err := config.ReadMysqlConfig(flags.DbConfig)

		if err != nil {
			utils.Colorize(utils.ColorRed, err.Error())
			return
		}

		db, err := database.ConnectMysql(config)
		if err != nil {
			log.Fatal(err)
		}

		err = readcsv.ReadExcelAndInsertData(db, flags.FileName,config.Table)
		if err != nil {
			utils.Colorize(utils.ColorRed, err.Error())
			return
		}

		defer db.Close()
	} else {
		utils.Colorize(utils.ColorRed, "database name should be `postgres` OR `mysql`")
		return
	}
}

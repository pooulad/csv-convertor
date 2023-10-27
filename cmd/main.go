package main

import (
	"fmt"

	"github.com/pooulad/csv-convertor/config"
	"github.com/pooulad/csv-convertor/internal/readflag"
	"github.com/pooulad/csv-convertor/utils"
)

func main() {
	flags, err := readflag.ReadFlag()
	if err != nil {
		utils.Colorize(utils.ColorRed, err.Error())
	}
	fmt.Println(flags)
	var dbConfig interface{}
	if flags.DbType == "postgres" {
		config, err := config.ReadPostgresConfig(flags.DbConfig)
		if err != nil {
			fmt.Println(err)
			return
		}
		dbConfig = config
	} else {
		config, err := config.ReadMysqlConfig(flags.DbConfig)
		if err != nil {
			fmt.Println(err)
			return
		}
		dbConfig = config
	}


	fmt.Printf("this is config : %s\n",dbConfig)
}

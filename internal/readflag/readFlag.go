package readflag

import (
	"errors"
	"flag"
	"fmt"
)

type FlagReturns struct {
	DbConfig string
	DbType   string
	FileName string
}

func ReadFlag() (*FlagReturns, error) {
	fr := FlagReturns{}
	flag.StringVar(&fr.FileName, "file", "./username.csv", "CSV file path")
	flag.StringVar(&fr.DbConfig, "config", "", "Config file path")
	flag.StringVar(&fr.DbType, "type", "", "Database type")

	flag.Parse()

	fmt.Println(fr.DbConfig)
	if fr.FileName == "" {
		return nil, errors.New("please enter a valid file address")
	}
	if fr.DbConfig == "" {
		return nil, errors.New("please enter a valid config address")
	}
	if fr.DbType == "" {
		return nil, errors.New("please enter a valid database name(postgres OR mysql)")
	}

	return &fr, nil
}

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

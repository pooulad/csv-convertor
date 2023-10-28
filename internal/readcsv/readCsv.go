package readcsv

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
)

type ListRecord struct {
	Username   string
	Identifier string
	Firstname  string
	Lastname   string
}

func ReadExcelAndInsertData(db *sql.DB, fileAddress string, table_name string) error {
	f, err := os.Open(fileAddress)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	csvData := createList(data)
	for v, r := range csvData {
		fmt.Println(v)
		for j, g := range r {
			fmt.Println(j)
			fmt.Println(g)
		}
		fmt.Println(r)
	}
	fmt.Println(csvData)
	return nil
}

func createList(data interface{}) [][]string {
	fmt.Println(data)
	return [][]string{{"ff", "f"}, {"dd", "d"}}
}

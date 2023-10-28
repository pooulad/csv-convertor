package readcsv

import (
	"context"
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

type Test struct {
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

	// csvData := createList(data)
	for v, r := range data {
		fmt.Println(v)
		for j, g := range r {
			fmt.Println(j)
			fmt.Println(g)
		}
		fmt.Println(r)
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}


    // a := make(map[string]interface{})
	test := Test{}
	n := data[0]
	fmt.Println(data[0])
	for _, row := range n {
		// test.(string(row)) = row
		test.(row) = row 
	}
	fmt.Println(n[0])


	for i, row := range data {
		for j, g := range row {

			// cellValue1 := row.Username
			// cellValue2 := row.Identifier
			// cellValue3 := row.Firstname
			// cellValue4 := row.Lastname

			s := fmt.Sprintf("INSERT INTO users (Username,Identifier,Firstname,Lastname) VALUES ($1,$2,$3,$4)")

			_, err = tx.Exec(
				"INSERT INTO users (Username,Identifier,Firstname,Lastname) VALUES ($1,$2,$3,$4)",
				cellValue1,
				cellValue2,
				cellValue3,
				cellValue4,
			)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	// fmt.Println(csvData)
	return nil
}

func createList(data interface{}) [][]string {
	fmt.Println(data)
	return [][]string{{"ff", "f"}, {"dd", "d"}}
}

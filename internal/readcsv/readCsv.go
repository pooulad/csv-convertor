package readcsv

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/pooulad/csv-convertor/utils"
)

func ReadExcelAndInsertData(db *sql.DB, fileAddress string, tableName string) error {
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

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	columnRow := data[0]
	columnString := strings.Join(columnRow, ", ")
	for i, row := range data {
		if i != 0 {
			var quotedValues []string
			for _, value := range row {
				quotedValues = append(quotedValues, fmt.Sprintf("'%s'", value))
			}
			dataString := strings.Join(quotedValues, ", ")

			query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columnString, dataString)
			querySuccessfulMessage := fmt.Sprintf("[%s] inserted into %s", dataString, tableName)

			utils.Colorize(utils.ColorGreen, querySuccessfulMessage)

			_, err = tx.Exec(query)
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

	return nil
}

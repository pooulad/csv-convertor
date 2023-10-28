package readcsv

import (
	// "context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
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

	fmt.Println(csvData)

	// stmt, err := db.Prepare("INSERT INTO table_name (field1) VALUES (?)")
	// if err != nil {
	// 	// handle the error
	// }
	// defer stmt.Close()
	// ----------------------
	// tx, err := db.BeginTx(context.Background(), nil)
	// if err != nil {
	// 	return err
	// }

	// for _, row := range csvData {
	// 	cellValue1 := row.Username
	// 	cellValue2 := row.Identifier
	// 	cellValue3 := row.Firstname
	// 	cellValue4 := row.Lastname

	// 	_, err = tx.Exec(
	// 		"INSERT INTO users (Username,Identifier,Firstname,Lastname) VALUES ($1,$2,$3,$4)",
	// 		cellValue1,
	// 		cellValue2,
	// 		cellValue3,
	// 		cellValue4,
	// 	)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// }

	// err = tx.Commit()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func createList(data [][]string) []interface{} {
	var finalList []interface{}
	for i, line := range data {
		var rec interface{}
		if i == 0 {
			continue
		} else {
			dataType := reflect.TypeOf(line)
			dataValue := reflect.ValueOf(line)

			if dataValue.Kind() == reflect.Slice {
				for i := 0; i < dataValue.Len(); i++ {
					field := dataType.Field(i)
					fieldValue := dataValue.Field(i)
	

					fmt.Println(field)
					fmt.Println(fieldValue)
				}
			}
			fmt.Println(dataValue)
		}
		finalList = append(finalList, rec)
	}
	return finalList
}

// dataType := reflect.TypeOf(data)
// dataValue := reflect.ValueOf(data)

// fmt.Println(dataType)
// fmt.Println(dataValue)
// // var finalStruct []interface{}
// if dataType.Kind() == reflect.Struct {
// 	for i := 0; i < dataType.NumField(); i++ {
// 		field := dataType.Field(i)
// 		fieldValue := dataValue.Field(i)

// 		fmt.Printf("Field name: %s\n", field.Name)
// 		fmt.Printf("Field type: %v\n", field.Type)
// 		fmt.Printf("Field value: %v\n", fieldValue.Interface())
// 	}
// 	}else{

// 		fmt.Println("Field value")
// }

// // var finalList interface{}
// // for i, line := range data {
// // 	// var rec ListRecord
// // 	if i == 0 {
// // 		continue
// // 	} else {
// // 		for j, field := range line {
// // 			switch j {
// // 			case 0:
// // 				rec.Username = field
// // 			case 1:
// // 				rec.Identifier = field
// // 			case 2:
// // 				rec.Firstname = field
// // 			case 3:
// // 				rec.Lastname = field
// // 			}
// // 		}
// // 	}
// // 	finalList = append(finalList, rec)
// // }
// return nil

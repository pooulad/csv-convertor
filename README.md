## Introduction👨‍💻

🚨This program is written with GO to convert and insert your csv file in database(my-sql OR postgresql) 

<a href="https://www.coffeebede.com/poulad"><img class="img-fluid" src="https://coffeebede.ir/DashboardTemplateV2/app-assets/images/banner/default-yellow.svg" /></a>
## Usage

**flags**
```
🚨databaseType = "mysql" OR "postgres"
-type=databaseType or --type=databaseType
```
```
-config=config.json or --config=config.json
```
```
-file=file.csv or --file.csv=file.csv
```
## Config file is a json file that includes db information for insert csv

mysql sample : 
```json
{
  "username": "postgres",
  "password": "postgres",
  "host": "localhost",
  "name": "exel2sqldb",
  "port": "3306"
}
```

postgres sample : 
```json
{
  "username": "postgres",
  "password": "postgres",
  "host": "localhost",
  "name": "exel2sqldb",
  "ssl": "disable"
}
```

## Examples

I have prepared some examples for you that you can use

## Windows
```
go run main.go -type=postgres -file=./assets/username.csv -config=./postgres.json
```
OR
```
.\build.exe -type=postgres -file=./assets/username.csv -config=./postgres.json
```

## Linux
```
go run ./cmd/main.go -type=postgres -file=./assets/username.csv -config=./postgres.json
```
OR
```
sudo ./buildFileName ./cmd/main.go -type=postgres -file=./assets/username.csv -config=./postgres.json
```
OR

**you can use wine package too:**
```
wine start ./buildFileName -type=postgres -file=./assets/username.csv -config=./postgres.json
```


## Tech Stack

**Database:** postgresql,mysql

**Lang:** GO

    github.com/lib/pq
    github.com/go-sql-driver/mysql


## Authors

- [@pooulad](https://www.github.com/pooulad)
  
## Demo

![Screenshot 2023-10-24 211725](https://github.com/pooulad/csv-2-sql-example/assets/86445458/57dd39e1-9929-4a8c-8223-6f83167d23de)


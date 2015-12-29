// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// started from https://github.com/asdf072/struct-create

package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/exec"
	"strings"
)

var defaults = Configuration{
	DbUser:     os.Args[1],
	DbPassword: os.Args[2],
	DbName:     os.Args[3],
	PkgName:    "todb",
	TagLabel:   "db",
}

var config Configuration

type Configuration struct {
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
	// PkgName gives name of the package using the stucts
	PkgName string `json:"pkg_name"`
	// TagLabel produces tags commonly used to match database field names with Go struct members
	TagLabel string `json:"tag_label"`
}

type ColumnSchema struct {
	TableName              string
	ColumnName             string
	IsNullable             string
	DataType               string
	CharacterMaximumLength sql.NullInt64
	NumericPrecision       sql.NullInt64
	NumericScale           sql.NullInt64
	ColumnType             string
	ColumnKey              string
}

func writeFile(schemas []ColumnSchema, table string) (int, error) {
	file, err := os.Create("./generated/" + table + ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	header := "package " + config.PkgName + "\n\n"
	header += "import (\n"
	header += "\"fmt\"\n"

	sString := structString(schemas, table)

	if strings.Contains(sString, "null.") {
		header += "\"gopkg.in/guregu/null.v3\"\n"
	}
	if strings.Contains(sString, "time.") {
		header += "\"time\"\n"
	}
	header += ")\n\n"

	hString := handleString(table)
	totalBytes, err := fmt.Fprint(file, header+sString+hString)
	if err != nil {
		log.Fatal(err)
	}
	return totalBytes, nil
}

func handleString(table string) string {
	// out := "func handle" + formatName(table) + "()([]" + formatName(table) + ", error) {\n"
	out := "func handle" + formatName(table) + "(method string, id int)(interface{}, error) {\n"
	out += "    if method == \"GET\" {\n"
	out += "	    ret := []" + formatName(table) + "{}\n"
	out += "        if id >= 0 {\n"
	out += "    	    err := globalDB.Select(&ret, \"select * from " + table + " where id=?\", id)\n"
	out += "    	    if err != nil {\n"
	out += "    		    fmt.Println(err)\n"
	out += "    		    return nil, err\n"
	out += "    		}\n"
	out += "    	} else {\n"
	out += "    		queryStr := \"select * from " + table + "\"\n"
	out += "    	    err := globalDB.Select(&ret, queryStr)\n"
	out += "    	    if err != nil {\n"
	out += "    		    fmt.Println(err)\n"
	out += "    		    return nil, err\n"
	out += "    	    }\n"
	out += "   	    }\n"
	out += "    	return ret, nil\n"
	out += "    }\n"
	out += "    return nil, nil\n"
	out += "}\n\n"

	return out
}

func structString(schemas []ColumnSchema, table string) string {

	out := "type " + formatName(table) + " struct{\n"
	for _, cs := range schemas {

		if cs.TableName == table {

			goType, _, err := goType(&cs)

			if err != nil {
				log.Fatal(err)
			}
			out = out + "\t" + formatName(cs.ColumnName) + " " + goType
			if len(config.TagLabel) > 0 {
				out = out + "\t`" + config.TagLabel + ":\"" + cs.ColumnName + "\" json:\"" + formatNameLower(cs.ColumnName) + "\"`"
			}
			out = out + "\n"
		}
	}
	out = out + "}\n\n"

	return out
}

func getSchema() ([]ColumnSchema, []string) {
	conn, err := sql.Open("mysql", config.DbUser+":"+config.DbPassword+"@/information_schema")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	q := "SELECT TABLE_NAME, COLUMN_NAME, IS_NULLABLE, DATA_TYPE, " +
		"CHARACTER_MAXIMUM_LENGTH, NUMERIC_PRECISION, NUMERIC_SCALE, COLUMN_TYPE, " +
		"COLUMN_KEY FROM COLUMNS WHERE TABLE_SCHEMA = ? ORDER BY TABLE_NAME, ORDINAL_POSITION"
	rows, err := conn.Query(q, config.DbName)
	if err != nil {
		log.Fatal(err)
	}
	columns := []ColumnSchema{}
	for rows.Next() {
		cs := ColumnSchema{}
		err := rows.Scan(&cs.TableName, &cs.ColumnName, &cs.IsNullable, &cs.DataType,
			&cs.CharacterMaximumLength, &cs.NumericPrecision, &cs.NumericScale,
			&cs.ColumnType, &cs.ColumnKey)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, cs)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	q = "select TABLE_NAME from tables WHERE TABLE_SCHEMA = ? AND table_type='BASE TABLE'"
	rows, err = conn.Query(q, config.DbName)
	if err != nil {
		log.Fatal(err)
	}
	tables := []string{}
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		tables = append(tables, tableName)
	}
	return columns, tables
}

func formatName(name string) string {
	parts := strings.Split(name, "_")
	newName := ""
	for _, p := range parts {
		if len(p) < 1 {
			continue
		}
		newName = newName + strings.Replace(p, string(p[0]), strings.ToUpper(string(p[0])), 1)
	}
	return newName
}

func formatNameLower(name string) string {
	newName := formatName(name)
	newName = strings.Replace(newName, string(newName[0]), strings.ToLower(string(newName[0])), 1)
	return newName
}

func goType(col *ColumnSchema) (string, string, error) {
	requiredImport := ""
	if col.IsNullable == "YES" {
		requiredImport = "database/sql"
	}
	var gt string = ""
	switch col.DataType {
	case "char", "varchar", "enum", "text", "longtext", "mediumtext", "tinytext":
		if col.IsNullable == "YES" {
			gt = "null.String"
		} else {
			gt = "string"
		}
	case "blob", "mediumblob", "longblob", "varbinary", "binary":
		gt = "[]byte"
	case "date", "time", "datetime", "timestamp":
		gt, requiredImport = "time.Time", "time"
	case "tinyint", "smallint", "int", "mediumint", "bigint":
		if col.IsNullable == "YES" {
			gt = "null.Int"
		} else {
			gt = "int64"
		}
	case "float", "decimal", "double":
		if col.IsNullable == "YES" {
			gt = "null.Float"
		} else {
			gt = "float64"
		}
	}
	if gt == "" {
		n := col.TableName + "." + col.ColumnName
		return "", "", errors.New("No compatible datatype (" + col.DataType + ") for " + n + " found")
	}
	return gt, requiredImport, nil
}

func main() {

	config = defaults

	columns, tables := getSchema()
	fmt.Println(tables)
	for _, table := range tables {
		bytes, err := writeFile(columns, table)
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("go", "fmt", "./generated/"+table+".go")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: Ok %d\n", table, bytes)
	}
}
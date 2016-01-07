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

package api

import (
	"encoding/json"
	"errors"
	"github.com/Comcast/traffic_control/traffic_ops/goto2/db"
	"reflect"
	"time"
)

// genericGetById gets a generic value from SQL from a table via 'select where id ='
// The returnType MUST be a nil pointer to the type to be populated from SQL.
func genericGetById(id int, table string, returnTypeNilPtr interface{}) (interface{}, error) {
	nstmt, err := db.GlobalDB.PrepareNamed("select * from " + table + " where id=:id")
	if err != nil {
		return nil, err
	}
	defer nstmt.Close()

	returnType := reflect.TypeOf(returnTypeNilPtr).Elem()
	slicePtr := reflect.New(reflect.SliceOf(returnType))
	//	slicePtr.Elem().Set(reflect.MakeSlice(reflect.SliceOf(myType), 0, 0))
	err = nstmt.Select(slicePtr.Interface(), struct{ Id int64 }{int64(id)})
	if err != nil {
		return nil, err
	}

	return reflect.Indirect(slicePtr).Interface(), nil
}

func genericGet(table string, returnTypeNilPtr interface{}) (interface{}, error) {
	queryStr := "select * from " + table

	returnType := reflect.TypeOf(returnTypeNilPtr).Elem()
	slicePtr := reflect.New(reflect.SliceOf(returnType))
	err := db.GlobalDB.Select(slicePtr.Interface(), queryStr)
	if err != nil {
		return nil, err
	}
	return reflect.Indirect(slicePtr).Interface(), nil
}

// buildInsertQuery builds an insert query from obj's "db" tags (as used by sqlx)
// returns an error if obj is not a struct
// TODO(require mysql to enable ANSI_QUOTES and use SQL92 quotes for tables and fields?)
// TODO(put skipped fields [id, last_updated] as constant map somewhere? Take as parameter?)
// TODO(Take ID field name as parameter?)
// TODO(put in DB package?)
func buildInsertQuery(table string, obj interface{}) (string, error) {
	tp := reflect.TypeOf(obj)
	if tp.Kind() != reflect.Struct {
		return "", errors.New("object is not a struct: " + tp.Kind().String())
	}

	s := `INSERT INTO ` + table + `(`
	valueStr := ""
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		dbFieldName := f.Tag.Get("db")

		// skip id, last_updated, as done by the generate tool
		if dbFieldName == "" || dbFieldName == "id" || dbFieldName == "last_updated" {
			continue
		}

		s += dbFieldName + ", "
		valueStr += ":" + dbFieldName + ", "
	}

	if valueStr == "" {
		return "", errors.New("object has no 'db' tagged fields")
	}

	s = s[:len(s)-2]                      // remove trailing ,
	valueStr = valueStr[:len(valueStr)-2] // remove trailing ,
	s += ") VALUES (" + valueStr + ")"
	return s, nil
}

// buildUpdateQuery builds an update query from obj's "db" tags (as used by sqlx)
// returns an error if obj is not a struct
// TODO(require mysql to enable ANSI_QUOTES and use SQL92 quotes for tables and fields?)
// TODO(put skipped fields [id, last_updated] as constant map somewhere? Take as parameter?)
// TODO(put in DB package?)
func buildUpdateQuery(table string, obj interface{}) (string, error) {
	tp := reflect.TypeOf(obj)
	if tp.Kind() != reflect.Struct {
		return "", errors.New("object is not a struct: " + tp.Kind().String())
	}

	s := `UPDATE ` + table + ` SET `
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		dbFieldName := f.Tag.Get("db")

		// skip id, last_updated, as done by the generate tool
		if dbFieldName == "" || dbFieldName == "id" || dbFieldName == "last_updated" {
			continue
		}

		s += dbFieldName + " = :" + dbFieldName + ", "
	}

	s = s[:len(s)-2] // remove trailing ,
	s += " WHERE id = :id"
	return s, nil
}

func genericPost(payload []byte, table string, payloadTypeNilPtr interface{}) (interface{}, error) {
	tp := reflect.TypeOf(payloadTypeNilPtr).Elem()
	reflectPtr := reflect.New(tp)
	v := reflectPtr.Interface()

	err := json.Unmarshal(payload, &v)
	if err != nil {
		return nil, err
	}

	sqlStr, err := buildInsertQuery(table, reflect.Zero(tp).Interface())
	if err != nil {
		return nil, err
	}

	result, err := db.GlobalDB.NamedExec(sqlStr, v)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func genericPut(id int, payload []byte, table string, payloadTypeNilPtr interface{}) (interface{}, error) {
	tp := reflect.TypeOf(payloadTypeNilPtr).Elem()
	reflectPtr := reflect.New(tp)
	reflectVal := reflect.Indirect(reflectPtr)
	v := reflectPtr.Interface()

	err := json.Unmarshal(payload, &v)
	if err != nil {
		return nil, err
	}

	idFieldName := "Id" // TODO(Constant somewhere? Function parameter?)
	idFieldValue := reflectVal.FieldByName(idFieldName)
	emptyval := reflect.Value{}
	if idFieldValue == emptyval {
		return nil, errors.New("payload type has no Id field")
	}
	if !idFieldValue.CanSet() {
		return nil, errors.New("payload type Id field cannot be set.")
	}
	if idFieldValue.Kind() != reflect.Int64 {
		return nil, errors.New("payload type Id field is not an int64")
	}
	idFieldValue.SetInt(int64(id))

	lastUpdatedFieldName := "LastUpdated" // TODO(Constant somewhere? Function parameter?)
	lastUpdatedFieldValue := reflectVal.FieldByName(lastUpdatedFieldName)
	if lastUpdatedFieldValue != (reflect.Value{}) && lastUpdatedFieldValue.CanSet() && lastUpdatedFieldValue.Type() == reflect.TypeOf((*time.Time)(nil)) {
		lastUpdatedFieldValue.Set(reflect.ValueOf(time.Now()))
	}

	sqlString, err := buildUpdateQuery(table, reflect.Zero(tp).Interface())
	if err != nil {
		return nil, err
	}

	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		return nil, err
	}
	return result, err
}

func genericDelete(id int, table string) (interface{}, error) {
	result, err := db.GlobalDB.NamedExec("DELETE FROM "+table+" WHERE id=:id", struct{ Id int64 }{int64(id)})
	if err != nil {
		return nil, err
	}
	return result, err
}

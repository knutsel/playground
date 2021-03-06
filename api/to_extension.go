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

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type ToExtension struct {
	Id                    int64            `db:"id" json:"id"`
	Name                  string           `db:"name" json:"name"`
	Version               string           `db:"version" json:"version"`
	InfoUrl               string           `db:"info_url" json:"infoUrl"`
	ScriptFile            string           `db:"script_file" json:"scriptFile"`
	Isactive              int64            `db:"isactive" json:"isactive"`
	AdditionalConfigJson  null.String      `db:"additional_config_json" json:"additionalConfigJson"`
	Description           null.String      `db:"description" json:"description"`
	ServercheckShortName  null.String      `db:"servercheck_short_name" json:"servercheckShortName"`
	ServercheckColumnName null.String      `db:"servercheck_column_name" json:"servercheckColumnName"`
	LastUpdated           time.Time        `db:"last_updated" json:"lastUpdated"`
	Links                 ToExtensionLinks `json:"_links" db:-`
}

type ToExtensionLinks struct {
	Self     string   `db:"self" json:"_self"`
	TypeLink TypeLink `json:"type" db:-`
}

// @Title getToExtensionById
// @Description retrieves the to_extension information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    ToExtension
// @Resource /api/2.0
// @Router /api/2.0/to_extension/{id} [get]
func getToExtensionById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []ToExtension{}
	arg := ToExtension{}
	arg.Id = int64(id)
	queryStr := "select *, concat('" + API_PATH + "to_extension/', id) as self "
	queryStr += ", concat('" + API_PATH + "type/', type) as type_id_ref"
	queryStr += " from to_extension where id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getToExtensions
// @Description retrieves the to_extension
// @Accept  application/json
// @Success 200 {array}    ToExtension
// @Resource /api/2.0
// @Router /api/2.0/to_extension [get]
func getToExtensions(db *sqlx.DB) (interface{}, error) {
	ret := []ToExtension{}
	queryStr := "select *, concat('" + API_PATH + "to_extension/', id) as self "
	queryStr += ", concat('" + API_PATH + "type/', type) as type_id_ref"
	queryStr += " from to_extension"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postToExtension
// @Description enter a new to_extension
// @Accept  application/json
// @Param                 Body body     ToExtension   true "ToExtension object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/to_extension [post]
func postToExtension(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v ToExtension
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO to_extension("
	sqlString += "name"
	sqlString += ",version"
	sqlString += ",info_url"
	sqlString += ",script_file"
	sqlString += ",isactive"
	sqlString += ",additional_config_json"
	sqlString += ",description"
	sqlString += ",servercheck_short_name"
	sqlString += ",servercheck_column_name"
	sqlString += ",type"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:version"
	sqlString += ",:info_url"
	sqlString += ",:script_file"
	sqlString += ",:isactive"
	sqlString += ",:additional_config_json"
	sqlString += ",:description"
	sqlString += ",:servercheck_short_name"
	sqlString += ",:servercheck_column_name"
	sqlString += ",:type"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putToExtension
// @Description modify an existing to_extensionentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     ToExtension   true "ToExtension object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/to_extension/{id}  [put]
func putToExtension(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v ToExtension
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE to_extension SET "
	sqlString += "name = :name"
	sqlString += ",version = :version"
	sqlString += ",info_url = :info_url"
	sqlString += ",script_file = :script_file"
	sqlString += ",isactive = :isactive"
	sqlString += ",additional_config_json = :additional_config_json"
	sqlString += ",description = :description"
	sqlString += ",servercheck_short_name = :servercheck_short_name"
	sqlString += ",servercheck_column_name = :servercheck_column_name"
	sqlString += ",type = :type"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delToExtensionById
// @Description deletes to_extension information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    ToExtension
// @Resource /api/2.0
// @Router /api/2.0/to_extension/{id} [delete]
func delToExtension(id int, db *sqlx.DB) (interface{}, error) {
	arg := ToExtension{}
	arg.Id = int64(id)
	result, err := db.NamedExec("DELETE FROM to_extension WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

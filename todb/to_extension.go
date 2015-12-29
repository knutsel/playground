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

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package todb

import (
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"time"
)

type ToExtension struct {
	Id                    int64       `db:"id" json:"id"`
	Name                  string      `db:"name" json:"name"`
	Version               string      `db:"version" json:"version"`
	InfoUrl               string      `db:"info_url" json:"infoUrl"`
	ScriptFile            string      `db:"script_file" json:"scriptFile"`
	Isactive              int64       `db:"isactive" json:"isactive"`
	AdditionalConfigJson  null.String `db:"additional_config_json" json:"additionalConfigJson"`
	Description           null.String `db:"description" json:"description"`
	ServercheckShortName  null.String `db:"servercheck_short_name" json:"servercheckShortName"`
	ServercheckColumnName null.String `db:"servercheck_column_name" json:"servercheckColumnName"`
	Type                  int64       `db:"type" json:"type"`
	LastUpdated           time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleToExtension(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		ret := []ToExtension{}
		if id >= 0 {
			err := globalDB.Select(&ret, "select * from to_extension where id=$1", id)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		} else {
			queryStr := "select * from to_extension"
			err := globalDB.Select(&ret, queryStr)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		}
		return ret, nil
	} else if method == "POST" {
		var v Asn
		err := json.Unmarshal(payload, &v)
		if err != nil {
			fmt.Println(err)
		}
		insertString := "INSERT INTO to_extension("
		insertString += "name"
		insertString += ",version"
		insertString += ",info_url"
		insertString += ",script_file"
		insertString += ",isactive"
		insertString += ",additional_config_json"
		insertString += ",description"
		insertString += ",servercheck_short_name"
		insertString += ",servercheck_column_name"
		insertString += ",type"
		insertString += ") VALUES ("
		insertString += ":name"
		insertString += ",:version"
		insertString += ",:info_url"
		insertString += ",:script_file"
		insertString += ",:isactive"
		insertString += ",:additional_config_json"
		insertString += ",:description"
		insertString += ",:servercheck_short_name"
		insertString += ",:servercheck_column_name"
		insertString += ",:type"
		insertString += ")"
		result, err := globalDB.NamedExec(insertString, v)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return result.LastInsertId()
	}
	return nil, nil
}
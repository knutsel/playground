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

type Federation struct {
	Id          int64       `db:"id" json:"id"`
	Cname       string      `db:"cname" json:"cname"`
	Description null.String `db:"description" json:"description"`
	Ttl         int64       `db:"ttl" json:"ttl"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleFederation(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		ret := []Federation{}
		if id >= 0 {
			err := globalDB.Select(&ret, "select * from federation where id=$1", id)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		} else {
			queryStr := "select * from federation"
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
		insertString := "INSERT INTO federation("
		insertString += "cname"
		insertString += ",description"
		insertString += ",ttl"
		insertString += ") VALUES ("
		insertString += ":cname"
		insertString += ",:description"
		insertString += ",:ttl"
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
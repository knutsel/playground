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

package api

import (
	_ "github.com/Comcast/traffic_control/traffic_ops/goto2/output_format" // needed for swagger
	"time"
)

type Regex struct {
	Id          int64     `db:"id" json:"id"`
	Pattern     string    `db:"pattern" json:"pattern"`
	Type        int64     `db:"type" json:"type"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleRegex(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getRegex(id)
	} else if method == "POST" {
		return postRegex(payload)
	} else if method == "PUT" {
		return putRegex(id, payload)
	} else if method == "DELETE" {
		return delRegex(id)
	}
	return nil, nil
}

func getRegex(id int) (interface{}, error) {
	if id >= 0 {
		return getRegexById(id)
	} else {
		return getRegexs()
	}
}

// @Title getRegexById
// @Description retrieves the regex information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Regex
// @Resource /api/2.0
// @Router /api/2.0/regex/{id} [get]
func getRegexById(id int) (interface{}, error) {
	return genericGetById(id, "regex", (*Regex)(nil))
}

// @Title getRegexs
// @Description retrieves the regex information for a certain id
// @Accept  application/json
// @Success 200 {array}    Regex
// @Resource /api/2.0
// @Router /api/2.0/regex [get]
func getRegexs() (interface{}, error) {
	return genericGet("regex", (*Regex)(nil))
}

// @Title postRegex
// @Description enter a new regex
// @Accept  application/json
// @Param              Pattern json     string   false "pattern description"
// @Param                 Type json      int64   false "type description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/regex [post]
func postRegex(payload []byte) (interface{}, error) {
	return genericPost(payload, "regex", (*Regex)(nil))
}

// @Title putRegex
// @Description modify an existing regexentry
// @Accept  application/json
// @Param              Pattern json     string   false "pattern description"
// @Param                 Type json      int64   false "type description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/regex [put]
func putRegex(id int, payload []byte) (interface{}, error) {
	return genericPut(id, payload, "regex", (*Regex)(nil))
}

// @Title delRegexById
// @Description deletes regex information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Regex
// @Resource /api/2.0
// @Router /api/2.0/regex/{id} [delete]
func delRegex(id int) (interface{}, error) {
	return genericDelete(id, "regex")
}

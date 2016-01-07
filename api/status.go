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
	null "gopkg.in/guregu/null.v3"
	"time"
)

type Status struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	Description null.String `db:"description" json:"description"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleStatus(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getStatus(id)
	} else if method == "POST" {
		return postStatus(payload)
	} else if method == "PUT" {
		return putStatus(id, payload)
	} else if method == "DELETE" {
		return delStatus(id)
	}
	return nil, nil
}

func getStatus(id int) (interface{}, error) {
	if id >= 0 {
		return getStatusById(id)
	} else {
		return getStatuss()
	}
}

// @Title getStatusById
// @Description retrieves the status information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Status
// @Resource /api/2.0
// @Router /api/2.0/status/{id} [get]
func getStatusById(id int) (interface{}, error) {
	return genericGetById(id, "status", (*Status)(nil))
}

// @Title getStatuss
// @Description retrieves the status information for a certain id
// @Accept  application/json
// @Success 200 {array}    Status
// @Resource /api/2.0
// @Router /api/2.0/status [get]
func getStatuss() (interface{}, error) {
	return genericGet("status", (*Status)(nil))
}

// @Title postStatus
// @Description enter a new status
// @Accept  application/json
// @Param                 Name json     string   false "name description"
// @Param          Description json     string    true "description description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/status [post]
func postStatus(payload []byte) (interface{}, error) {
	return genericPost(payload, "status", (*Status)(nil))
}

// @Title putStatus
// @Description modify an existing statusentry
// @Accept  application/json
// @Param                 Name json     string   false "name description"
// @Param          Description json     string    true "description description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/status [put]
func putStatus(id int, payload []byte) (interface{}, error) {
	return genericPut(id, payload, "status", (*Status)(nil))
}

// @Title delStatusById
// @Description deletes status information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Status
// @Resource /api/2.0
// @Router /api/2.0/status/{id} [delete]
func delStatus(id int) (interface{}, error) {
	return genericDelete(id, "status")
}

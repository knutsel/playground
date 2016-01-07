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

type Cdn struct {
	Id          int64       `db:"id" json:"id"`
	Name        null.String `db:"name" json:"name"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleCdn(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getCdn(id)
	} else if method == "POST" {
		return postCdn(payload)
	} else if method == "PUT" {
		return putCdn(id, payload)
	} else if method == "DELETE" {
		return delCdn(id)
	}
	return nil, nil
}

func getCdn(id int) (interface{}, error) {
	if id >= 0 {
		return getCdnById(id)
	} else {
		return getCdns()
	}
}

// @Title getCdnById
// @Description retrieves the cdn information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Cdn
// @Resource /api/2.0
// @Router /api/2.0/cdn/{id} [get]
func getCdnById(id int) (interface{}, error) {
	return genericGetById(id, "cdn", (*Cdn)(nil))
}

// @Title getCdns
// @Description retrieves the cdn information for a certain id
// @Accept  application/json
// @Success 200 {array}    Cdn
// @Resource /api/2.0
// @Router /api/2.0/cdn [get]
func getCdns() (interface{}, error) {
	return genericGet("cdn", (*Cdn)(nil))
}

// @Title postCdn
// @Description enter a new cdn
// @Accept  application/json
// @Param                 Name json     string    true "name description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/cdn [post]
func postCdn(payload []byte) (interface{}, error) {
	return genericPost(payload, "cdn", (*Cdn)(nil))
}

// @Title putCdn
// @Description modify an existing cdnentry
// @Accept  application/json
// @Param                 Name json     string    true "name description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/cdn [put]
func putCdn(id int, payload []byte) (interface{}, error) {
	return genericPut(id, payload, "cdn", (*Cdn)(nil))
}

// @Title delCdnById
// @Description deletes cdn information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Cdn
// @Resource /api/2.0
// @Router /api/2.0/cdn/{id} [delete]
func delCdn(id int) (interface{}, error) {
	return genericDelete(id, "cdn")
}

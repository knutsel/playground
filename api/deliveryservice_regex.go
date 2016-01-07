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
)

type DeliveryserviceRegex struct {
	Deliveryservice int64    `db:"deliveryservice" json:"deliveryservice"`
	Regex           int64    `db:"regex" json:"regex"`
	SetNumber       null.Int `db:"set_number" json:"setNumber"`
}

func handleDeliveryserviceRegex(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getDeliveryserviceRegex(id)
	} else if method == "POST" {
		return postDeliveryserviceRegex(payload)
	} else if method == "PUT" {
		return putDeliveryserviceRegex(id, payload)
	} else if method == "DELETE" {
		return delDeliveryserviceRegex(id)
	}
	return nil, nil
}

func getDeliveryserviceRegex(id int) (interface{}, error) {
	if id >= 0 {
		return getDeliveryserviceRegexById(id)
	} else {
		return getDeliveryserviceRegexs()
	}
}

// @Title getDeliveryserviceRegexById
// @Description retrieves the deliveryservice_regex information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    DeliveryserviceRegex
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_regex/{id} [get]
func getDeliveryserviceRegexById(id int) (interface{}, error) {
	return genericGetById(id, "deliveryservice_regex", (*DeliveryserviceRegex)(nil))
}

// @Title getDeliveryserviceRegexs
// @Description retrieves the deliveryservice_regex information for a certain id
// @Accept  application/json
// @Success 200 {array}    DeliveryserviceRegex
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_regex [get]
func getDeliveryserviceRegexs() (interface{}, error) {
	return genericGet("deliveryservice_regex", (*DeliveryserviceRegex)(nil))
}

// @Title postDeliveryserviceRegex
// @Description enter a new deliveryservice_regex
// @Accept  application/json
// @Param      Deliveryservice json      int64   false "deliveryservice description"
// @Param                Regex json      int64   false "regex description"
// @Param            SetNumber json        int    true "set_number description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_regex [post]
func postDeliveryserviceRegex(payload []byte) (interface{}, error) {
	return genericPost(payload, "deliveryservice_regex", (*DeliveryserviceRegex)(nil))
}

// @Title putDeliveryserviceRegex
// @Description modify an existing deliveryservice_regexentry
// @Accept  application/json
// @Param      Deliveryservice json      int64   false "deliveryservice description"
// @Param                Regex json      int64   false "regex description"
// @Param            SetNumber json        int    true "set_number description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_regex [put]
func putDeliveryserviceRegex(id int, payload []byte) (interface{}, error) {
	return genericPut(id, payload, "deliveryservice_regex", (*DeliveryserviceRegex)(nil))
}

// @Title delDeliveryserviceRegexById
// @Description deletes deliveryservice_regex information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    DeliveryserviceRegex
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice_regex/{id} [delete]
func delDeliveryserviceRegex(id int) (interface{}, error) {
	return genericDelete(id, "deliveryservice_regex")
}

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

type PhysLocation struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	ShortName   string      `db:"short_name" json:"shortName"`
	Address     string      `db:"address" json:"address"`
	City        string      `db:"city" json:"city"`
	State       string      `db:"state" json:"state"`
	Zip         string      `db:"zip" json:"zip"`
	Poc         null.String `db:"poc" json:"poc"`
	Phone       null.String `db:"phone" json:"phone"`
	Email       null.String `db:"email" json:"email"`
	Comments    null.String `db:"comments" json:"comments"`
	Region      int64       `db:"region" json:"region"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handlePhysLocation(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getPhysLocation(id)
	} else if method == "POST" {
		return postPhysLocation(payload)
	} else if method == "PUT" {
		return putPhysLocation(id, payload)
	} else if method == "DELETE" {
		return delPhysLocation(id)
	}
	return nil, nil
}

func getPhysLocation(id int) (interface{}, error) {
	if id >= 0 {
		return getPhysLocationById(id)
	} else {
		return getPhysLocations()
	}
}

// @Title getPhysLocationById
// @Description retrieves the phys_location information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    PhysLocation
// @Resource /api/2.0
// @Router /api/2.0/phys_location/{id} [get]
func getPhysLocationById(id int) (interface{}, error) {
	return genericGetById(id, "phys_location", (*PhysLocation)(nil))
}

// @Title getPhysLocations
// @Description retrieves the phys_location information for a certain id
// @Accept  application/json
// @Success 200 {array}    PhysLocation
// @Resource /api/2.0
// @Router /api/2.0/phys_location [get]
func getPhysLocations() (interface{}, error) {
	return genericGet("phys_location", (*PhysLocation)(nil))
}

// @Title postPhysLocation
// @Description enter a new phys_location
// @Accept  application/json
// @Param                 Name json     string   false "name description"
// @Param            ShortName json     string   false "short_name description"
// @Param              Address json     string   false "address description"
// @Param                 City json     string   false "city description"
// @Param                State json     string   false "state description"
// @Param                  Zip json     string   false "zip description"
// @Param                  Poc json     string    true "poc description"
// @Param                Phone json     string    true "phone description"
// @Param                Email json     string    true "email description"
// @Param             Comments json     string    true "comments description"
// @Param               Region json      int64   false "region description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/phys_location [post]
func postPhysLocation(payload []byte) (interface{}, error) {
	return genericPost(payload, "phys_location", (*PhysLocation)(nil))
}

// @Title putPhysLocation
// @Description modify an existing phys_locationentry
// @Accept  application/json
// @Param                 Name json     string   false "name description"
// @Param            ShortName json     string   false "short_name description"
// @Param              Address json     string   false "address description"
// @Param                 City json     string   false "city description"
// @Param                State json     string   false "state description"
// @Param                  Zip json     string   false "zip description"
// @Param                  Poc json     string    true "poc description"
// @Param                Phone json     string    true "phone description"
// @Param                Email json     string    true "email description"
// @Param             Comments json     string    true "comments description"
// @Param               Region json      int64   false "region description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/phys_location [put]
func putPhysLocation(id int, payload []byte) (interface{}, error) {
	return genericPut(id, payload, "phys_location", (*PhysLocation)(nil))
}

// @Title delPhysLocationById
// @Description deletes phys_location information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    PhysLocation
// @Resource /api/2.0
// @Router /api/2.0/phys_location/{id} [delete]
func delPhysLocation(id int) (interface{}, error) {
	return genericDelete(id, "phys_location")
}

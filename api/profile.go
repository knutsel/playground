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

type Profile struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	Description null.String `db:"description" json:"description"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleProfile(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getProfile(id)
	} else if method == "POST" {
		return postProfile(payload)
	} else if method == "PUT" {
		return putProfile(id, payload)
	} else if method == "DELETE" {
		return delProfile(id)
	}
	return nil, nil
}

func getProfile(id int) (interface{}, error) {
	if id >= 0 {
		return getProfileById(id)
	} else {
		return getProfiles()
	}
}

// @Title getProfileById
// @Description retrieves the profile information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Profile
// @Resource /api/2.0
// @Router /api/2.0/profile/{id} [get]
func getProfileById(id int) (interface{}, error) {
	return genericGetById(id, "profile", (*Profile)(nil))
}

// @Title getProfiles
// @Description retrieves the profile information for a certain id
// @Accept  application/json
// @Success 200 {array}    Profile
// @Resource /api/2.0
// @Router /api/2.0/profile [get]
func getProfiles() (interface{}, error) {
	return genericGet("profile", (*Profile)(nil))
}

// @Title postProfile
// @Description enter a new profile
// @Accept  application/json
// @Param                 Name json     string   false "name description"
// @Param          Description json     string    true "description description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/profile [post]
func postProfile(payload []byte) (interface{}, error) {
	return genericPost(payload, "profile", (*Profile)(nil))
}

// @Title putProfile
// @Description modify an existing profileentry
// @Accept  application/json
// @Param                 Name json     string   false "name description"
// @Param          Description json     string    true "description description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/profile [put]
func putProfile(id int, payload []byte) (interface{}, error) {
	return genericPut(id, payload, "profile", (*Profile)(nil))
}

// @Title delProfileById
// @Description deletes profile information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Profile
// @Resource /api/2.0
// @Router /api/2.0/profile/{id} [delete]
func delProfile(id int) (interface{}, error) {
	return genericDelete(id, "profile")
}

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

type CachegroupParameter struct {
	Cachegroup  int64     `db:"cachegroup" json:"cachegroup"`
	Parameter   int64     `db:"parameter" json:"parameter"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleCachegroupParameter(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getCachegroupParameter(id)
	} else if method == "POST" {
		return postCachegroupParameter(payload)
	} else if method == "PUT" {
		return putCachegroupParameter(id, payload)
	} else if method == "DELETE" {
		return delCachegroupParameter(id)
	}
	return nil, nil
}

func getCachegroupParameter(id int) (interface{}, error) {
	if id >= 0 {
		return getCachegroupParameterById(id)
	} else {
		return getCachegroupParameters()
	}
}

// @Title getCachegroupParameterById
// @Description retrieves the cachegroup_parameter information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    CachegroupParameter
// @Resource /api/2.0
// @Router /api/2.0/cachegroup_parameter/{id} [get]
func getCachegroupParameterById(id int) (interface{}, error) {
	return genericGetById(id, "cachegroup_parameter", (*CachegroupParameter)(nil))
}

// @Title getCachegroupParameters
// @Description retrieves the cachegroup_parameter information for a certain id
// @Accept  application/json
// @Success 200 {array}    CachegroupParameter
// @Resource /api/2.0
// @Router /api/2.0/cachegroup_parameter [get]
func getCachegroupParameters() (interface{}, error) {
	return genericGet("cachegroup_parameter", (*CachegroupParameter)(nil))
}

// @Title postCachegroupParameter
// @Description enter a new cachegroup_parameter
// @Accept  application/json
// @Param           Cachegroup json      int64   false "cachegroup description"
// @Param            Parameter json      int64   false "parameter description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/cachegroup_parameter [post]
func postCachegroupParameter(payload []byte) (interface{}, error) {
	return genericPost(payload, "cachegroup_parameter", (*CachegroupParameter)(nil))
}

// @Title putCachegroupParameter
// @Description modify an existing cachegroup_parameterentry
// @Accept  application/json
// @Param           Cachegroup json      int64   false "cachegroup description"
// @Param            Parameter json      int64   false "parameter description"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/cachegroup_parameter [put]
func putCachegroupParameter(id int, payload []byte) (interface{}, error) {
	return genericPut(id, payload, "cachegroup_parameter", (*CachegroupParameter)(nil))
}

// @Title delCachegroupParameterById
// @Description deletes cachegroup_parameter information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    CachegroupParameter
// @Resource /api/2.0
// @Router /api/2.0/cachegroup_parameter/{id} [delete]
func delCachegroupParameter(id int) (interface{}, error) {
	return genericDelete(id, "cachegroup_parameter")
}

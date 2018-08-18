/* mi_schoolDistrict_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_schoolDistrict_database
 */

package controllers

import (
    "encoding/json"
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/services"
)

type SchoolDistrictController struct {
    schoolDistrictService services.ISchoolDistrictService
}
func NewSchoolDistrictController(schoolDistrictService services.ISchoolDistrictService) *SchoolDistrictController {
    return &SchoolDistrictController{
        schoolDistrictService: schoolDistrictService,
    }
}
func (p *SchoolDistrictController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    queryStr := r.URL.Query().Get("query")
    query := skaioskit.QueryRequest{}
    err := json.Unmarshal([]byte(queryStr), &query)

    if err != nil {
        return skaioskit.ControllerResponse{Status: http.StatusBadRequest, Body: skaioskit.EmptyResponse{}}
    }

    schoolDistricts, count, err := p.schoolDistrictService.GetSchoolDistricts(query)

    if err == nil {
        return skaioskit.ControllerResponse{Status: http.StatusOK, Body: GetSchoolDistrictsResponse{SchoolDistricts: schoolDistricts, Total: count}}
    } else {
        panic(err)
    }
}
func (p *SchoolDistrictController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *SchoolDistrictController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *SchoolDistrictController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}

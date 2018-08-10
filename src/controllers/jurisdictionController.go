/* mi_jurisdiction_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_jurisdiction_database
 */

package controllers

import (
    "encoding/json"
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/core"
    "skaioskit/services"
)

type JurisdictionController struct {
    jurisdictionService services.IJurisdictionService
}
func NewJurisdictionController(jurisdictionService services.IJurisdictionService) *JurisdictionController {
    return &JurisdictionController{
        jurisdictionService: jurisdictionService,
    }
}
func (p *JurisdictionController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    queryStr := r.URL.Query().Get("query")
    query := core.QueryRequest{}
    err := json.Unmarshal([]byte(queryStr), &query)

    if err != nil {
        return skaioskit.ControllerResponse{Status: http.StatusBadRequest, Body: skaioskit.EmptyResponse{}}
    }

    jurisdictions, count, err := p.jurisdictionService.GetJurisdictions(query)

    if err == nil {
        return skaioskit.ControllerResponse{Status: http.StatusOK, Body: GetJurisdictionsResponse{Jurisdictions: jurisdictions, Total: count}}
    } else {
        panic(err)
    }
}
func (p *JurisdictionController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *JurisdictionController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *JurisdictionController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}

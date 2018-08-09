package controllers

import (
    "encoding/json"
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/core"
    "skaioskit/services"
)

type VoterController struct {
    voterService services.IVoterService
}
func NewVoterController(voterService services.IVoterService) *VoterController {
    return &VoterController{
        voterService: voterService,
    }
}
func (p *VoterController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    queryStr := r.URL.Query().Get("query")
    query := core.QueryRequest{}
    err := json.Unmarshal([]byte(queryStr), &query)
    if err != nil {
        return skaioskit.ControllerResponse{Status: http.StatusBadRequest, Body: skaioskit.EmptyResponse{}}
    }

    voters, err := p.voterService.GetVoters(query)
    if err == nil {
        return skaioskit.ControllerResponse{Status: http.StatusOK, Body: core.GetVotersResponse{Voters: voters}}
    } else {
        return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
    }
}
func (p *VoterController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *VoterController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    count := p.voterService.GetVoterCount()
    return skaioskit.ControllerResponse{Status: http.StatusOK, Body: core.GetVoterCount{Count: count}}
}
func (p *VoterController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}

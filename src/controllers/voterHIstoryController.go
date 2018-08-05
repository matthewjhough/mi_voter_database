package controllers

import (
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/core"
    "skaioskit/services"
)

type VoterHistoryController struct {
    voterHistoryService services.IVoterHistoryService
}
func NewVoterHistoryController(voterHistoryService services.IVoterHistoryService) *VoterHistoryController {
    return &VoterHistoryController{
        voterHistoryService: voterHistoryService,
    }
}
func (p *VoterHistoryController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *VoterHistoryController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *VoterHistoryController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    count := p.voterHistoryService.GetVoterHistoryCount()
    return skaioskit.ControllerResponse{Status: http.StatusOK, Body: core.GetVoterHistoryCount{Count: count}}
}
func (p *VoterHistoryController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}

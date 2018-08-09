package controllers

import (
    "skaioskit/models"
)

type GetVotersResponse struct {
    Voters []models.Voter
}

type GetAboutResponse struct {
    Version string
    BuildTime string
}

/* mi_voter_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_voter_database
 */

package controllers

import (
    "skaioskit/models"
)

//TODO: These should probably be defined in the same go file as the controllers than use them.

type GetAboutResponse struct {
    CoreVersion string
    Version string
    BuildTime string
}

type GetCountiesResponse struct {
    Counties []models.County
    Total uint64
}

type GetElectionsResponse struct {
    Elections []models.Election
    Total uint64
}

type GetJurisdictionsResponse struct {
    Jurisdictions []models.Jurisdiction
    Total uint64
}

type GetSchoolDistrictsResponse struct {
    SchoolDistricts []models.SchoolDistrict
    Total uint64
}

type GetVillagesResponse struct {
    Villages []models.Village
    Total uint64
}

type GetVotersResponse struct {
    Voters []models.Voter
    Total uint64
}

type GetVoterHistoriesResponse struct {
    VoterHistories []models.VoterHistory
    Total uint64
}

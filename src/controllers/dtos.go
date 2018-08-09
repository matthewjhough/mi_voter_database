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

type GetVotersResponse struct {
    Voters []models.Voter
}

type GetAboutResponse struct {
    CoreVersion string
    Version string
    BuildTime string
}

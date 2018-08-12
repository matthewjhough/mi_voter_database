/* mi_voter_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_voter_database
 */

package providers

import (
    "skaioskit/models"
)

type IVoterDataProvider interface {
    ParseCounties() <-chan models.County
    ParseJurisdictions() <-chan models.Jurisdiction
    ParseSchools() <-chan models.SchoolDistrict
    ParseVillages() <-chan models.Village
    ParseElections() <-chan models.Election
    ParseVoters() <-chan models.Voter
    ParseVoterHistories() <-chan models.VoterHistory
}

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
    "skaioskit/services"
)

type IVoterDataProvider interface {
    EnsureCounties(services.ICountyService)
    EnsureJurisdictions(services.IJurisdictionService)
    EnsureSchools(services.ISchoolDistrictService)
    EnsureVillages(services.IVillageService)
    EnsureElections(services.IElectionService)
    EnsureVoters(services.IVoterService)
    EnsureVoterHistories(services.IVoterHistoryService)
}

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

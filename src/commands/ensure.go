/* mi_voter_database
 * 
 * Copyright (C) 2018 Nathan Mentley - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the BSD license.
 *
 * You should have received a copy of the BSD license with
 * this file. If not, please visit: https://github.com/nathanmentley/mi_voter_database
 */

package commands

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/services"
    "skaioskit/providers"
)

var ensureCmd = &cobra.Command{
    Use:   "ensure",
    Short: "imports the database",
    Long:  `ensures the database schema exists and has imported the voter data.`,
    Run: func(cmd *cobra.Command, args []string) {
        //setup db connection
        db, err := gorm.Open("mysql", viper.GetString("mysql-connection-str"))
        if err != nil {
            panic(err)
        }
        defer db.Close()

        //setup services
        schoolService := services.NewSchoolDistrictService(db)
        countyService := services.NewCountyService(db)
        villageService := services.NewVillageService(db)
        jurisdictionService := services.NewJurisdictionService(db)
        electionService := services.NewElectionService(db)
        voterService := services.NewVoterService(db)
        voterHistoryService := services.NewVoterHistoryService(db)

        provider := providers.NewMichiganByteWidthDataProvider()
        //ensure db
        countyService.EnsureCountyTable()
        for county := range provider.ParseCounties() {
            countyService.EnsureCounty(county)
        }
        jurisdictionService.EnsureJurisdictionTable()
        for jurisdiction := range provider.ParseJurisdictions() {
            jurisdictionService.EnsureJurisdiction(jurisdiction)
        }
        schoolService.EnsureSchoolDistrictTable()
        for school := range provider.ParseSchools() {
            schoolService.EnsureSchoolDistrict(school)
        }
        villageService.EnsureVillageTable()
        for village := range provider.ParseVillages() {
            villageService.EnsureVillage(village)
        }
        electionService.EnsureElectionTable()
        for election := range provider.ParseElections() {
            electionService.EnsureElection(election)
        }
        voterService.EnsureVoterTable()
        for voter := range provider.ParseVoters() {
            voterService.EnsureVoter(voter)
        }
        voterHistoryService.EnsureVoterHistoryTable()
        for voterHistory := range provider.ParseVoterHistories() {
            voterHistoryService.EnsureVoterHistory(voterHistory)
        }
    },
}

//Entry
func init() {
    RootCmd.AddCommand(ensureCmd)
}

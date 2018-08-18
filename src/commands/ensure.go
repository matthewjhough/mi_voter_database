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

//The ensure command will load data from an IVoterDataProvider and ensure each record exists and is updated in the database.
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
        // TODO: Issue#12 For sanity we should setup a DI pattern or something to setup instances of all of this.
        //                  That would make unit testing a bit more sane and we'd be able to probably call common code
        //                  to get all these services instead of setting up each one at the start of a command.
        schoolService := services.NewSchoolDistrictService(db)
        countyService := services.NewCountyService(db)
        villageService := services.NewVillageService(db)
        jurisdictionService := services.NewJurisdictionService(db)
        electionService := services.NewElectionService(db)
        voterService := services.NewVoterService(db)
        voterHistoryService := services.NewVoterHistoryService(db)

        // TODO: Issue#8 - We should get the provider dynamically based on the args.
        //                  This will be needed to support multiple states worth of voter data.
        //                  We could also use this to support pulling data from multiple sources for the same state.
        provider := providers.NewMichiganByteWidthDataProvider()

        //ensure db

        //setup records for all the counties, jurisdictions, schools, and villages. These will be referenced by the voter records.
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

        //ensure every election refereced by the provider has a record.
        electionService.EnsureElectionTable()
        for election := range provider.ParseElections() {
            electionService.EnsureElection(election)
        }

        //ensure we have a record for each voter and records for each election they voted in.
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

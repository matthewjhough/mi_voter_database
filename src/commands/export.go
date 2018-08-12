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

    //"skaioskit/providers"
)

var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Generates csv files from data provider",
    Long:  `Generates csv files from the configured data provider / data set.`,
    Run: func(cmd *cobra.Command, args []string) {
        /*
        provider := providers.NewMichiganByteWidthDataProvider()

        for county := range provider.ParseCounties() {
        }
        for jurisdiction := range provider.ParseJurisdictions() {
        }
        for school := range provider.ParseSchools() {
        }
        for village := range provider.ParseVillages() {
        }
        for election := range provider.ParseElections() {
        }
        for voter := range provider.ParseVoters() {
        }
        for voterHistory := range provider.ParseVoterHistories() {
        }
        */
    },
}

//Entry
func init() {
    RootCmd.AddCommand(exportCmd)
}

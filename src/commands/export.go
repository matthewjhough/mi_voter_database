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
    "os"
    "encoding/csv"

    "github.com/spf13/cobra"
    "github.com/jinzhu/copier"

    "skaioskit/models"
    "skaioskit/providers"
)

var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Generates csv files from data provider",
    Long:  `Generates csv files from the configured data provider / data set.`,
    Run: func(cmd *cobra.Command, args []string) {
        provider := providers.NewMichiganByteWidthDataProvider()

        writeCounties(provider)
        /*
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

func writeCounties(provider providers.IVoterDataProvider) {
    chnl := make(chan models.IExportable)
    go func() {
        for county := range provider.ParseCounties() {
            obj := models.County{}
            copier.Copy(&obj, &county)
            chnl <- &obj
        }
        close(chnl)
    }()
    writeCsv("/working/export/counties.csv", models.GetCountyCSVHeader(), chnl)
}

func writeCsv(filename string, header []string, chnl <-chan models.IExportable) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    w := csv.NewWriter(file)
    w.Write(header)
    for record := range chnl {
        if err := w.Write(record.ToSlice()); err != nil {
            panic(err)
        }
    }
    w.Flush()
}

//Entry
func init() {
    RootCmd.AddCommand(exportCmd)
}

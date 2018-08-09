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
    "net/http"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/services"
    "skaioskit/controllers"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "runs the rest api",
    Long:  `runs the rest api that exposes the voter / voter history data for other microservices to consume.`,
    Run: func(cmd *cobra.Command, args []string) {
        //setup db connection
        db, err := gorm.Open("mysql", viper.GetString("mysql-connection-str"))
        if err != nil {
            panic(err)
        }
        defer db.Close()

        //setup services
        //schoolService := services.NewSchoolDistrictService(db)
        //countyService := services.NewCountyService(db)
        //jurisdictionService := services.NewJurisdictionService(db)
        //electionService := services.NewElectionService(db)
        voterService := services.NewVoterService(db)
        voterHistoryService := services.NewVoterHistoryService(db)

        //build controllers
        aboutController := skaioskit.NewControllerProcessor(controllers.NewAboutController())
        voterController := skaioskit.NewControllerProcessor(controllers.NewVoterController(voterService))
        voterHistoryController := skaioskit.NewControllerProcessor(controllers.NewVoterHistoryController(voterHistoryService))

        //setup routing to controllers
        r := mux.NewRouter()
        r.HandleFunc("/about", aboutController.Logic)
        r.HandleFunc("/voter", voterController.Logic)
        r.HandleFunc("/voterHistory", voterHistoryController.Logic)

        //wrap everything behind a jwt middleware
        jwtMiddleware := skaioskit.JWTEnforceMiddleware([]byte(viper.GetString("jwt-key")))
        http.Handle("/", skaioskit.PanicHandler(jwtMiddleware(r)))

        //server up app
        if err := http.ListenAndServe(":" + viper.GetString("port-number"), nil); err != nil {
            panic(err)
        }
    },
}

//Entry
func init() {
    RootCmd.AddCommand(serveCmd)
}

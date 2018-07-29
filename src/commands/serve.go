package commands

import (
    "os"
    "net/http"

    "github.com/spf13/cobra"
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
        db, err := gorm.Open("mysql", os.Getenv("APP_MYSQL_CONN_STR"))
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

        //build controllers
        voterController := skaioskit.NewControllerProcessor(controllers.NewVoterController(voterService))

        //setup routing to controllers
        r := mux.NewRouter()
        r.HandleFunc("/voter", voterController.Logic)

        //wrap everything behind a jwt middleware
        jwtMiddleware := skaioskit.JWTEnforceMiddleware([]byte(os.Getenv("APP_JWT_KEY")))
        http.Handle("/", jwtMiddleware(r))

        //server up app
        if err := http.ListenAndServe(":" + os.Getenv("APP_PORT_NUMBER"), nil); err != nil {
            panic(err)
        }
    },
}

//Entry
func init() {
    RootCmd.AddCommand(serveCmd)
}

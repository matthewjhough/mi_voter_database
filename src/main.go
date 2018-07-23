package main

import (
    "os"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"
)

//Entry
func main() {
    //setup db connection
    db, err := gorm.Open("mysql", os.Getenv("APP_MYSQL_CONN_STR"))
    if err != nil {
        panic(err)
    }
    defer db.Close()

    //build controllers
    voterController := skaioskit.NewControllerProcessor(NewVoterController())

    //setup routing to controllers
    //auth end points
    jwtMiddleware := skaioskit.JWTEnforceMiddleware([]byte(os.Getenv("APP_JWT_KEY")))
    r := mux.NewRouter()
    r.HandleFunc("/voter", voterController.Logic)

    http.Handle("/", jwtMiddleware(r))

    //server up app
    if err := http.ListenAndServe(":" + os.Getenv("APP_PORT_NUMBER"), nil); err != nil {
        panic(err)
    }
}

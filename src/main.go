package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"
)

func ensureSchools(service ISchoolDistrictService) {
    schools := []SchoolDistrict{}

    file, err := os.Open("/data/schoolcd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        code, err := strconv.ParseUint(strings.TrimLeft(line[0:12], "0"), 0, 64)
        if err != nil {
            panic(err)
        }

        school := SchoolDistrict{Code: code, Name: line[12:]}
        schools = append(schools, school)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureSchoolDistricts(schools)
}

func ensureCounties(service ICountyService) {
    counties := []County{}

    file, err := os.Open("/data/countycd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        code, err := strconv.ParseUint(strings.TrimLeft(line[0:2], "0"), 0, 32)
        if err != nil {
            panic(err)
        }

        county := County{Code: uint(code), Name: line[2:]}
        counties = append(counties, county)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureCounties(counties)
}

func ensureJurisdictions(service IJurisdictionService) {
    counties := []Jurisdiction{}

    file, err := os.Open("/data/jurisdcd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        code, err := strconv.ParseUint(strings.TrimLeft(line[0:7], "0"), 0, 32)
        if err != nil {
            panic(err)
        }

        county := Jurisdiction{Code: uint(code), Name: line[7:]}
        counties = append(counties, county)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureJurisdictions(counties)
}

func ensureElections(service IElectionService) {
    elections := []Election{}

    file, err := os.Open("/data/electionscd.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        code, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[0:21], " "), "0"), 0, 64)
        if err != nil {
            panic(err)
        }

        election := Election{Code: code, Name: line[21:]}
        elections = append(elections, election)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureElections(elections)
}

func ensureVoters(service IVoterService) {
    voters := []Voter{}

    file, err := os.Open("/data/entire_state_v.lst")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        voterId, err := strconv.ParseUint(strings.TrimLeft(strings.TrimLeft(line[448:461], " "), "0"), 0, 64)
        if err != nil {
            panic(err)
        }

        voter := Voter{
            VoterId: voterId,
            LastName: strings.Trim(line[0:35], " "),
            FirstName: strings.Trim(line[35:55], " "),
            MiddleName: strings.Trim(line[55:75], " "),
            NameSuffix: strings.Trim(line[75:78], " "),
            Gender: strings.Trim(line[82:83], " "),
        }
        voters = append(voters, voter)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureVoters(voters)
}

//Entry
func main() {
    //setup db connection
    db, err := gorm.Open("mysql", os.Getenv("APP_MYSQL_CONN_STR"))
    if err != nil {
        panic(err)
    }
    defer db.Close()

    //setup services
    schoolService := NewSchoolDistrictService(db)
    countyService := NewCountyService(db)
    jurisdictionService := NewJurisdictionService(db)
    electionService := NewElectionService(db)
    voterService := NewVoterService(db)

    //ensure db
    ensureSchools(schoolService)
    ensureCounties(countyService)
    ensureJurisdictions(jurisdictionService)
    ensureElections(electionService)
    ensureVoters(voterService)

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

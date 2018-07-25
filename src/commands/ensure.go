package commands

import (
    "os"
    "bufio"
    "strconv"
    "strings"

    "github.com/spf13/cobra"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "skaioskit/core"
    "skaioskit/services"
)

func ensureSchools(service services.ISchoolDistrictService) {
    schools := []core.SchoolDistrict{}

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

        school := core.SchoolDistrict{Code: code, Name: line[12:]}
        schools = append(schools, school)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureSchoolDistricts(schools)
}

func ensureCounties(service services.ICountyService) {
    counties := []core.County{}

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

        county := core.County{Code: uint(code), Name: line[2:]}
        counties = append(counties, county)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureCounties(counties)
}

func ensureJurisdictions(service services.IJurisdictionService) {
    counties := []core.Jurisdiction{}

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

        county := core.Jurisdiction{Code: uint(code), Name: line[7:]}
        counties = append(counties, county)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureJurisdictions(counties)
}

func ensureElections(service services.IElectionService) {
    elections := []core.Election{}

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

        election := core.Election{Code: code, Name: line[21:]}
        elections = append(elections, election)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    service.EnsureElections(elections)
}

func ensureVoters(db *gorm.DB, service services.IVoterService) {
    service.EnsureVoterTable()
    tx := db.Begin()
    counter := 0

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
            //panic(err)
            //TODO Log
        } else {
            counter++
            voter := core.Voter{
                VoterId: voterId,
                LastName: strings.Trim(line[0:35], " "),
                FirstName: strings.Trim(line[35:55], " "),
                MiddleName: strings.Trim(line[55:75], " "),
                NameSuffix: strings.Trim(line[75:78], " "),
                Gender: strings.Trim(line[82:83], " "),
            }
            tx.Create(&voter)

            if counter > 50000 {
                counter = 0
                tx.Commit()
                tx = db.Begin()
            }
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    tx.Commit()
}

var ensureCmd = &cobra.Command{
    Use:   "ensure",
    Short: "imports the database",
    Long:  `ensures the database schema exists and has imported the voter data.`,
    Run: func(cmd *cobra.Command, args []string) {
        //setup db connection
        db, err := gorm.Open("mysql", os.Getenv("APP_MYSQL_CONN_STR"))
        if err != nil {
            panic(err)
        }
        defer db.Close()

        //setup services
        schoolService := services.NewSchoolDistrictService(db)
        countyService := services.NewCountyService(db)
        jurisdictionService := services.NewJurisdictionService(db)
        electionService := services.NewElectionService(db)
        voterService := services.NewVoterService(db)

        //ensure db
        ensureSchools(schoolService)
        ensureCounties(countyService)
        ensureJurisdictions(jurisdictionService)
        ensureElections(electionService)
        ensureVoters(db, voterService)
    },
}

//Entry
func init() {
    RootCmd.AddCommand(ensureCmd)
}

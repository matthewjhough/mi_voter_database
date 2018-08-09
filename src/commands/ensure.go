package commands

import (
    "go.pedge.io/inject"
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

        module := inject.NewModule()
        module.BindInterface((*services.ISchoolDistrictService)(nil)).ToConstructor(services.NewSchoolDistrictService);

        //setup services
        schoolService := services.NewSchoolDistrictService(db)
        countyService := services.NewCountyService(db)
        villageService := services.NewVillageService(db)
        jurisdictionService := services.NewJurisdictionService(db)
        electionService := services.NewElectionService(db)
        voterService := services.NewVoterService(db)
        voterHistoryService := services.NewVoterHistoryService(db)

        provider := providers.NewMichiganByteWidthDataProvider(db)
        //ensure db
        provider.EnsureCounties(countyService)
        provider.EnsureJurisdictions(jurisdictionService)
        provider.EnsureSchools(schoolService)
        provider.EnsureVillages(villageService)
        provider.EnsureElections(electionService)
        provider.EnsureVoters(voterService)
        provider.EnsureVoterHistories(voterHistoryService)
    },
}

//Entry
func init() {
    RootCmd.AddCommand(ensureCmd)
}

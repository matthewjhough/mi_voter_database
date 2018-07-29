package commands

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "cobra-example",
    Short: "An example of cobra",
    Long: `This application shows how to create modern CLI 
applications in go using Cobra CLI library`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    //    Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    viper.SetConfigType("yaml")
    viper.AddConfigPath("/etc/skaioskit/")
    viper.SetConfigName("config.yaml")

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}

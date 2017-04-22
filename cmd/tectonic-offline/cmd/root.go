package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var tectonicImagesVar string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tectonic-offline",
	Short: "Manages container images required for Offline Tectonic Installation",
	Long:  `Longer`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().BoolP("help", "h", false, "Help message")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	if tectonicImagesVar != "" { 
		viper.SetConfigFile(tectonicImagesVar)
	}

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")               // adding home directory as first search path
	viper.AutomaticEnv()                       // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
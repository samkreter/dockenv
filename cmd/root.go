package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/samkreter/dockdev/util"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dockdev",
	Short: "dockdev provides easy way to setup isolated docker development enviorments.",
	Long: `dockdev uses the same notion of a kubernetes pod with the a pause contianer 
to create a local enviorment for development that doesn't use the hosts networking or IPC.`,
}

var create = &cobra.Command{
	Use:   "create",
	Short: "Creates a pod or single container.",
	Long: `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		err := util.Create()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var list = &cobra.Command{
	Use:   "list",
	Short: "List the current avalible dev enviorments.",
	Long: `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		err := util.List()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var clean = &cobra.Command{
	Use:   "clean",
	Short: "Stop and remove all current containers.",
	Long: `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		err := util.Clean()
		if err != nil {
			log.Fatal(err)
		}
	},
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.dockdev.yaml)")
	

	//Add the sub commands
	RootCmd.AddCommand(create)
	RootCmd.AddCommand(list)
	RootCmd.AddCommand(clean)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if config != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config)
	} else {
		// Search config in home directory with name ".dockdev" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dockdev")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

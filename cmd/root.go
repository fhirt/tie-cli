/*
Copyright © 2025 fhirt
*/
package cmd

import (
	"log"
	"os"

	"github.com/fhirt/tie-cli/cmd/auth"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile    string
	tieCliHome string
	rootCmd    = &cobra.Command{
		Use:   "tie-cli",
		Short: "tie-cli is a command line tool for various tie utilities",
		Long:  `tie-cli is a command line tool for various tie utilities.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	userHome, err := os.UserHomeDir()
	cobra.CheckErr(err)
	tieCliHome = userHome + string(os.PathSeparator) + ".tie-cli"
	// make dir if not exists
	err = os.Mkdir(tieCliHome, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	configFilePath := tieCliHome + string(os.PathSeparator) + "tie-cli.yaml"
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		file, err := os.Create(configFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tie-cli/tie-cli.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("tie-cli")  // name of config file (without extension)
		viper.AddConfigPath(".")        // add the working directory as
		viper.AddConfigPath(tieCliHome) // adding home directory as additional search path
		viper.SetConfigType("yaml")
	}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TIE_CLI")

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

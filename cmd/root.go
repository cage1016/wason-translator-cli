/*
Copyright © 2022 KAI CHU CHUNG cage.chung@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "document-translator-cli",
	Short: "Translate Document via IBM Cloud Language Translator",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.document-translator-cli.yaml)")
	rootCmd.PersistentFlags().StringP("api_key", "", "API KEY", "IBM Cloud Language Translator API KEY")
	rootCmd.PersistentFlags().StringP("url", "", "URL", "IBM Cloud Language Translator API URL")
	rootCmd.PersistentFlags().StringP("version", "", "VERSION", "IBM Cloud Language Translator API VERSION")
	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api_key"))
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("version"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".document-translator-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".document-translator-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

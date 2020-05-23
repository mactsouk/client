/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/mactsouk/handlers"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var UserData handlers.UserPass

var SERVER string
var PORT string
var USERNAME string
var PASSWORD string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "A client for a RESTful server",
	Long:  `This utility is a client for a RESTful server`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// CheckHTTPcode checks whether the two HTTP codes are the same
func CheckHTTPcode(wanted, given int) bool {
	log.Println("Comparing", wanted, given)
	return wanted == given
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.client.yaml)")
	rootCmd.PersistentFlags().StringVarP(&USERNAME, "username", "u", "", "Username issuing the command")
	rootCmd.PersistentFlags().StringVarP(&PASSWORD, "password", "p", "", "Password issuing the command")
	rootCmd.PersistentFlags().StringVar(&SERVER, "server", "http://localhost", "RESTful server hostname")
	rootCmd.PersistentFlags().StringVar(&PORT, "port", ":1234", "Port of RESTful Server")

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

		// Search config in home directory with name ".client" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".client")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

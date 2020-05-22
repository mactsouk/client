/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userinfoCmd represents the userinfo command
var userinfoCmd = &cobra.Command{
	Use:   "userinfo",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("userinfo called")
	},
}

func init() {
	rootCmd.AddCommand(userinfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userinfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

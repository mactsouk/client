/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("time called")
		UserData.Username = "admin"
		UserData.Password = "admin"
		fmt.Printf("*** #%v\n", UserData)
		err := UserData.Validate()
		if err != nil {
			fmt.Println("IsUserAdmin - Validate:", err)
			return
		}

		UserData.Username = ""
		fmt.Printf("#%v\n", UserData)
		err = UserData.Validate()
		if err != nil {
			fmt.Println("IsUserAdmin - Validate:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

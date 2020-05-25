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
	Run:   GetUserInfo,
}

func GetUserInfo(cmd *cobra.Command, args []string) {
	fmt.Println("userinfo called")
}

func init() {
	rootCmd.AddCommand(userinfoCmd)
}

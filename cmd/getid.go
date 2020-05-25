/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getidCmd represents the getid command
var getidCmd = &cobra.Command{
	Use:   "getid",
	Short: "A brief description of your command",
	Long:  ``,
	Run:   GetID,
}

func GetID(cmd *cobra.Command, args []string) {
	fmt.Println("getid called")
}

func init() {
	rootCmd.AddCommand(getidCmd)
}

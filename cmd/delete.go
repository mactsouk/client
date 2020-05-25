/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  ``,
	Run:   Delete,
}

func Delete(cmd *cobra.Command, args []string) {
	fmt.Println("delete called")
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

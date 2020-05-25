/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  ``,
	Run:   Add,
}

func Add(cmd *cobra.Command, args []string) {
	fmt.Println("add called")
}

func init() {
	rootCmd.AddCommand(addCmd)
}

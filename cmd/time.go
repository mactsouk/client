/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "For the /time endpoint",
	Long:  `Visiting the /time endpoint with GET HTTP request`,
	Run:   TimeFunction,
}

// TimeFunction implements the functionality of the time command
func TimeFunction(cmd *cobra.Command, args []string) {
	fmt.Println("time called")
	req, err := http.NewRequest("GET", SERVER+PORT+"/time", nil)
	if err != nil {
		fmt.Println("Error in req: ", err)
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		fmt.Println(resp)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(data)
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

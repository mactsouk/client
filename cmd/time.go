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
	req, err := http.NewRequest("GET", SERVER+PORT+"/time", nil)
	if err != nil {
		fmt.Println("Timefunction – Error in req: ", err)
		return
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		fmt.Println(resp)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(data))
}

func init() {
	rootCmd.AddCommand(timeCmd)
}

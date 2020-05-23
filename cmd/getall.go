/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mactsouk/handlers"
	"github.com/spf13/cobra"
)

// getallCmd represents the getall command
var getallCmd = &cobra.Command{
	Use:   "getall",
	Short: "A brief description of your command",
	Long:  ``,
	Run:   GetAll,
}

func GetAll(cmd *cobra.Command, args []string) {
	fmt.Println("getall called")
	req, err := http.NewRequest("GET", SERVER+PORT+"/getall", nil)
	if err != nil {
		fmt.Println("GetAll – Error in req: ", err)
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

	var users []handlers.User
	handlers.SliceFromJSON(users, resp.Body)
	data, err := handlers.PrettyJSON(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func init() {
	rootCmd.AddCommand(getallCmd)
}

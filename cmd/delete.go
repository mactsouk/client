/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mactsouk/handlers"
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
	USERID, _ := cmd.Flags().GetInt("id")
	if USERID <= 0 {
		fmt.Println("User ID value should be greater than 0!", USERID)
		return
	}

	fmt.Println("delete called")
	userpass := handlers.UserPass{Username: USERNAME, Password: PASSWORD}
	fmt.Println(userpass)

	// bytes.Buffer is both a Reader and a Writer
	buf := new(bytes.Buffer)
	err := userpass.ToJSON(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("DELETE", SERVER+PORT+"/username/"+strconv.Itoa(USERID), buf)
	if err != nil {
		fmt.Println("GetAll – Error in req: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Full response:", resp)
		fmt.Println("Response code:", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int("id", -1, "User ID")
}

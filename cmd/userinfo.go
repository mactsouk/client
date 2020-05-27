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

// userinfoCmd represents the userinfo command
var userinfoCmd = &cobra.Command{
	Use:   "userinfo",
	Short: "This command returns the user information given a user ID.",
	Long:  ``,
	Run:   GetUserInfo,
}

func GetUserInfo(cmd *cobra.Command, args []string) {
	USERID, _ := cmd.Flags().GetInt("id")
	if USERID <= 0 {
		fmt.Println("User ID value should be greater than 0!", USERID)
		return
	}

	userpass := handlers.UserPass{Username: USERNAME, Password: PASSWORD}

	// bytes.Buffer is both a Reader and a Writer
	buf := new(bytes.Buffer)
	err := userpass.ToJSON(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("GET", SERVER+PORT+"/username/"+strconv.Itoa(USERID), buf)
	if err != nil {
		fmt.Println("GetUserInfo – Error in req: ", err)
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

	var returnUser = handlers.User{}
	err = returnUser.FromJSON(resp.Body)
	if err != nil {
		fmt.Println("GetUserInfo:", err)
		return
	}

	t, err := handlers.PrettyJSON(returnUser)
	if err != nil {
		fmt.Println("PrettyJSON:", err)
		return
	}
	fmt.Println(t)
	defer resp.Body.Close()
}

func init() {
	rootCmd.AddCommand(userinfoCmd)
	userinfoCmd.Flags().Int("id", -1, "User ID")
}

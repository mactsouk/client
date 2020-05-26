/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mactsouk/handlers"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "This command is for updating existing users.",
	Long:  ``,
	Run:   Update,
}

func Update(cmd *cobra.Command, args []string) {
	userpass := handlers.Input{Username: USERNAME, Password: PASSWORD, Admin: 0}

	if len(load) == 0 {
		fmt.Println("No user data:", load)
		return
	}

	// Convert load into handlers.Input Structure
	var newUserData handlers.Input
	temp := []byte(load)
	err := json.Unmarshal(temp, &newUserData)
	if err != nil {
		fmt.Println("Update – error umarshalling user input:", load)
		return
	}

	fmt.Println("newUserData:", newUserData)
	userSlice := []handlers.Input{}
	userSlice = append(userSlice, userpass)
	userSlice = append(userSlice, newUserData)
	fmt.Println("userSlice:", userSlice)

	// bytes.Buffer is both a Reader and a Writer
	buf := new(bytes.Buffer)
	err = handlers.SliceToJSON(userSlice, buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("PUT", SERVER+PORT+"/update", buf)
	if err != nil {
		fmt.Println("GetID – Error in req: ", err)
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
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GetID:", err)
		return
	}

	fmt.Print(string(data))
	defer resp.Body.Close()
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&load, "data", "d", "", "User Record - {username, password, admin}")
}

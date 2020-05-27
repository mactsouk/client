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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding new users to the system.",
	Long:  `The input is in the handlers.Input format.`,
	Run:   Add,
}

func Add(cmd *cobra.Command, args []string) {
	userpass := handlers.Input{Username: USERNAME, Password: PASSWORD, Admin: 0}

	if len(DATA) == 0 {
		fmt.Println("No user data:", DATA)
		return
	}

	// Convert load into handlers.Input Structure
	var newUserData handlers.Input
	temp := []byte(DATA)
	err := json.Unmarshal(temp, &newUserData)
	if err != nil {
		fmt.Println("Add – error umarshalling user input:", DATA)
		return
	}

	userSlice := []handlers.Input{}
	userSlice = append(userSlice, userpass)
	userSlice = append(userSlice, newUserData)

	// bytes.Buffer is both a Reader and a Writer
	buf := new(bytes.Buffer)
	err = handlers.SliceToJSON(userSlice, buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", SERVER+PORT+"/add", buf)
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
	rootCmd.AddCommand(addCmd)
}

/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mactsouk/handlers"
	"github.com/spf13/cobra"
)

// getidCmd represents the getid command
var getidCmd = &cobra.Command{
	Use:   "getid",
	Short: "Get the ID of a user given a handlers.UserPass structure.",
	Long:  ``,
	Run:   GetID,
}

func GetID(cmd *cobra.Command, args []string) {
	userpass := handlers.UserPass{Username: USERNAME, Password: PASSWORD}

	// bytes.Buffer is both a Reader and a Writer
	buf := new(bytes.Buffer)
	err := userpass.ToJSON(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("GET", SERVER+PORT+"/getid", buf)
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
	rootCmd.AddCommand(getidCmd)
}

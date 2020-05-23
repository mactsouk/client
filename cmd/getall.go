/*
Copyright © 2020 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"bytes"
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
	userpass := handlers.UserPass{Username: USERNAME, Password: PASSWORD}
	fmt.Println(userpass)

	// bytes.Buffer is both a Reader and a Writer
	buf := new(bytes.Buffer)
	err := userpass.ToJSON(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("GET", SERVER+PORT+"/getall", buf)
	if err != nil {
		fmt.Println("GetAll – Error in req: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := c.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
		return
	}
	defer resp.Body.Close()

	var users = []handlers.User{}
	handlers.SliceFromJSON(users, resp.Body)
	data, err := handlers.PrettyJSON(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(data)
}

func init() {
	rootCmd.AddCommand(getallCmd)
}

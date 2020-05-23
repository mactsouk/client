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

// loggedCmd represents the logged command
var loggedCmd = &cobra.Command{
	Use:   "logged",
	Short: "A brief description of your command",
	Long:  ``,
	Run:   LoggedUsers,
}

func LoggedUsers(cmd *cobra.Command, args []string) {
	fmt.Println("logged called")
	userpass := handlers.UserPass{Username: USERNAME, Password: PASSWORD}
	fmt.Println(userpass)

	// bytes.Buffer is both a Reader and a Writer
	buf := new(bytes.Buffer)
	err := userpass.ToJSON(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("GET", SERVER+PORT+"/logged", buf)
	if err != nil {
		fmt.Println("LoggedUsers – Error in req: ", err)
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
	handlers.SliceFromJSON(&users, resp.Body)
	data, err := handlers.PrettyJSON(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Data: ", data)
}

func init() {
	rootCmd.AddCommand(loggedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loggedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loggedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

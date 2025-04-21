package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "Get activity of a user",
	Long:  "An easy way to get the github activity of a user",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		username := strings.Join(args, "")
		getGithubUserActivities(username)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func getGithubUserActivities(username string) {
	fmt.Printf("Your username %s", username)
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("user-activity.json", body, 0666)

	var activities []map[string]interface{}

	err = json.Unmarshal(body, &activities)
	if err != nil {
		log.Fatal(err)
	}
	// Loop over each activity
	for _, ac := range activities {
		fmt.Println("Type:", ac["type"])
		fmt.Println("Repo:", ac["repo"])
		fmt.Println("Created At:", ac["created_at"])
		fmt.Println("---")
	}
}

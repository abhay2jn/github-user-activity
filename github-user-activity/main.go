package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/spf13/cobra"
)

type GitHubActivity struct {
	Type   string `json:"type"`
	Repo   struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
}

func getUserActivity(username string) ([]GitHubActivity, error) {
	// GitHub API 
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var events []GitHubActivity
	if err := json.NewDecoder(res.Body).Decode(&events); err != nil {
		return nil, err
	}
	return events, nil
}

func main() {
	var username string

	var rootCmd = &cobra.Command{
		Use:   "github user activity",
		Short: "Fetching recent user GitHub activity",
		Run: func(cmd *cobra.Command, args []string) {
			events, err := getUserActivity(username)
			if err != nil {
				log.Fatalf("Error while fetching data: %v\n", err)
			}

			if len(events) == 0 {
				fmt.Println("User has no recent activity on GitHub")
			} else {
				fmt.Printf("%s recent Activity:\n", username)
				for _, e := range events {
					fmt.Printf("%s: %s at %s\n", e.CreatedAt, e.Type, e.Repo.Name)
				}
			}
		},
	}
	rootCmd.Flags().StringVarP(&username, "user", "u", "", "GitHub username to fetch activity for")
	rootCmd.MarkFlagRequired("user")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

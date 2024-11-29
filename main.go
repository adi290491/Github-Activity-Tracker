package main

import (
	"encoding/json"
	"fmt"
	"github-tracker/event"
	"log"
	"net/http"
)

func main() {
	log.Println("----Github Activity Tracker----")

	fmt.Print(">github-activity ")
	var username string
	fmt.Scanf("%s", &username)
	log.Println("username: ", username)
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	fmt.Println("\n---Recent Activity---")
	var events []event.Event

	err = json.NewDecoder(response.Body).Decode(&events)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("\n---Events---")
	for _, event := range events {
		fmt.Println(event)
		if event.Type == "PushEvent" {
			for _, commit := range event.Payload.Commits {
				fmt.Printf("Author: %s, Commit: %s\n", commit.Author.Name, commit.Message)
				fmt.Println()
			}
		}
		fmt.Println("---")
		fmt.Println()
	}

}

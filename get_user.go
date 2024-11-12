package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func getUserById(userId string) *slack.User {
    api := slack.New(slack_token())
    user, err := api.GetUserInfo(userId)
    if err != nil {
	    fmt.Printf("%s\n", err)
	    panic(err)
    }
	return user
}

func getUserByEmail(email string) *slack.User {
	api := slack.New(slack_token())
	user, err := api.GetUserByEmail(email)
	if err != nil {
	    fmt.Printf("%s\n", err)
	    panic(err)
	}
	return user
}
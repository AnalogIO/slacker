package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func getUserById(userId string) {
    api := slack.New(slack_token())
    user, err := api.GetUserInfo(userId)
    if err != nil {
	    fmt.Printf("%s\n", err)
	    return
    }
    fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}

func getUserByEmail(email string) {
	api := slack.New(slack_token())
	user, err := api.GetUserByEmail(email)
	if err != nil {
	    fmt.Printf("%s\n", err)
	    return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

const (
	c0 = "Morning"
	c1 = "10-12"
	c2 = "12-14"
	c3 = "Closing"
)

func getAllUserGroups() {
	api := slack.New(slack_token())
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New(slack_token(), slack.OptionDebug(true))
	groups, err := api.GetUserGroups(slack.GetUserGroupsOptionIncludeUsers(false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}
}

func createUserGroups(list emails) {
	api := slack.New(slack_token())
	for _, email := range list {
		_, err := api.CreateUserGroup(email, slack.UserGroupOptionHandle(email))
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
	}
}

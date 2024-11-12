package main

import "os"

func slack_token() string {
	return os.Getenv("SLACK_TOKEN")
}
package main

import "os"

func slack_token() string {
	if token, ok := os.LookupEnv("SLACK_TOKEN"); ok {
		return token
	}
	panic("SLACK_TOKEN environment variable not set")
}

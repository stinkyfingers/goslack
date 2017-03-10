package main

import (
	"flag"
	"log"

	"github.com/stinkyfingers/goslack/upload"
)

var (
	file    = flag.String("f", "", "file to post")
	channel = flag.String("c", "cloud-distro", "channel (or comma sep channels)")
	token   = flag.String("t", "", "oauth token")
)

func main() {
	flag.Parse()

	if *file == "" {
		log.Fatal("no file provided")
	}
	path := *file

	if *token == "" {
		log.Fatal("no token provided")
	}

	err := upload.Upload(path, *channel, *token)
	if err != nil {
		log.Fatal(err)
	}

}

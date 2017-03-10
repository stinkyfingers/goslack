package main

import (
	"flag"
	"log"
	"strings"

	"github.com/stinkyfingers/goslack/message"
	"github.com/stinkyfingers/goslack/upload"
)

var (
	file      = flag.String("f", "", "file to post")
	channel   = flag.String("c", "cloud-distro", "channel (or comma sep channels)")
	token     = flag.String("t", "", "oauth token")
	buildInfo = flag.String("b", "", "comma-delim list of build details")
	hook      = flag.String("h", "", "webhook")
)

func main() {
	flag.Parse()

	if *file != "" {
		if *token == "" {
			log.Fatal("no token provided")
		}
		path := *file

		err := upload.Upload(path, *channel, *token)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *buildInfo != "" {
		if *hook == "" {
			log.Fatal("no webhook provided")
		}
		arr := strings.Split(*buildInfo, ",")
		d := &message.Distro{}
		for _, s := range arr {
			pair := strings.Split(s, "=")
			if len(pair) != 2 {
				log.Print("odd pair: ", pair)
				continue
			}
			switch strings.ToLower(pair[0]) {
			case "installerbuild":
				d.InstallerBuild = pair[1]
			case "installerhash":
				d.InstallerHash = pair[1]
			case "installerversion":
				d.InstallerVersion = pair[1]
			case "tsversion":
				d.TSVersion = pair[1]
			case "sourcebuild":
				d.SourceBuild = pair[1]
			case "sourcehash":
				d.SourceHash = pair[1]
			case "sourceversion":
				d.SourceVersion = pair[1]
			}
		}
		err := d.Send(*hook, *channel)
		if err != nil {
			log.Fatal(err)
		}
	}

}

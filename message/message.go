package message

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Message struct {
	Channel  string `json:"channel"`
	Username string `json:"username,omitempty"`
	Text     string `json:"text"`
	Icon     string `json:"icon_emoji,omitempty"`
}

type Distro struct {
	InstallerBuild   string
	InstallerHash    string
	InstallerVersion string
	TSVersion        string
	SourceBuild      string
	SourceHash       string
	SourceVersion    string
}

func (d *Distro) Send(hook, channel string) error {
	u := fmt.Sprintf("https://hooks.slack.com/services/%s", hook)

	msg := &Message{
		Text:    d.Stringify(),
		Channel: channel,
	}

	js, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := http.PostForm(u, url.Values{
		"payload": {string(js)},
	})
	if err != nil {
		return err
	}
	res.Body.Close()
	return nil
}

func (d *Distro) Stringify() string {
	return fmt.Sprintf("Installer Build: %s\nInstallerHash: %s\nInstallerVersion: %s\nTSVersion: %s\nSourceBuild: %s\nSourceHash: %s\nInstallerVersion: %s\n",
		d.InstallerBuild, d.InstallerHash, d.InstallerVersion, d.TSVersion,
		d.SourceBuild, d.SourceHash, d.SourceVersion)
}

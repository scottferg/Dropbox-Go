package main

import (
	"fmt"
	"github.com/scottferg/Dropbox-Go/dropbox"
)

func main() {
	s := dropbox.Session{
		Oauth2AccessToken: "OAUTH_TOKEN",
		AccessType:        "app_folder",
	}
	delta, err := dropbox.GetDelta(s, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(delta)

}

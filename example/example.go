package main

import (
	"fmt"
	"github.com/scottferg/Dropbox-Go/api"
	"github.com/scottferg/Dropbox-Go/files"
	"github.com/scottferg/Dropbox-Go/session"
)

func main() {
	s := session.Session{
		AppKey:     "APP_KEY",
		AppSecret:  "APP_SECRET",
		AccessType: "app_folder",
		Token: session.AccessToken{
			Secret: "ACCESS_SECRET",
			Key:    "ACCESS_KEY",
		},
	}

	uriPath := api.Uri{
		Root: api.RootSandbox,
		Path: "NERDS/test_form.pdf",
	}

	// Upload a file
	if file, err := ioutil.ReadFile("./test_form.pdf"); err != nil {
		fmt.Println(err.Error())
	} else {
		m, err := files.UploadFile(s, file, uriPath, nil)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// Download the file
	if file, _, err := files.GetFile(s, uriPath, nil); err == nil {
		ioutil.WriteFile("./test_result.pdf", file, os.ModePerm)
	} else {
		fmt.Println(err.Error())
	}
}

package main

import (
	"fmt"
	"github.com/scottferg/Dropbox-Go/dropbox"
)

func main() {
	s := dropbox.Session{
		AppKey:     "APP_KEY",
		AppSecret:  "APP_SECRET",
		AccessType: "app_folder",
		Token: dropbox.AccessToken{
			Secret: "ACCESS_SECRET",
			Key:    "ACCESS_KEY",
		},
	}

	uriPath := dropbox.Uri{
		Root: dropbox.RootSandbox,
		Path: "NERDS/test_form.pdf",
	}

	// Upload a file
	if file, err := ioutil.ReadFile("./test_form.pdf"); err != nil {
		fmt.Println(err.Error())
	} else {
		m, err := dropbox.UploadFile(s, file, uriPath, nil)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// Download the file
	if file, _, err := dropbox.GetFile(s, uriPath, nil); err == nil {
		ioutil.WriteFile("./test_result.pdf", file, os.ModePerm)
	} else {
		fmt.Println(err.Error())
	}
}

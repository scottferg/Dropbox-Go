Dropbox-Go
==========

Go library for the Dropbox API

## Setup

        $ go get -u github.com/scottferg/Dropbox-Go/session
        $ go get -u github.com/scottferg/Dropbox-Go/account
        $ go get -u github.com/scottferg/Dropbox-Go/files
        $ go get -u github.com/scottferg/Dropbox-Go/fileops

## Authentication

The Dropbox API uses OAuth for it's authentication mechanism. You can authenticate
a session by using an instance of a Session interface:

        s := session.Session{
            AppKey: "yourappkey",
            AppSecret: "yourappsecret",
            AccessType: "app_folder",
        }

Once you have initialized a Session you can then obtain a request token:

        s.ObtainRequestToken()

You will need to authorize this token with the app, which will require visiting
a Dropbox URL. You can grab this URL with the following function:

        session.GenerateAuthorizeUrl(s.Token.Key)

Once you have authorized the request token for use, you can obtain the access token
much like you obtained the request token:

        s.ObtainAccessToken()

Your session is now authorized to make requests.

## Making Requests

All API methods take API parameters as parameters to the function. If a method isn't required simply
passing an empty string value will ignore it.

Dropbox-Go
==========

Go library for the Dropbox API

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

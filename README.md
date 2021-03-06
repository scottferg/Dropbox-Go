Dropbox-Go
==========

Go library for the Dropbox API

## Setup

        $ go get -u github.com/scottferg/Dropbox-Go/dropbox

## Authentication

The Dropbox API uses OAuth for it's authentication mechanism. You can authenticate
a session by using an instance of a Session interface:

        s := dropbox.Session{
            AppKey: "yourappkey",
            AppSecret: "yourappsecret",
            AccessType: "app_folder",
        }

Once you have initialized a Session you can then obtain a request token:

        s.ObtainRequestToken()

You will need to authorize this token with the app, which will require visiting
a Dropbox URL. You can grab this URL with the following function:

        dropbox.GenerateAuthorizeUrl(s.Token.Key)

Once you have authorized the request token for use, you can obtain the access token
much like you obtained the request token:

        s.ObtainAccessToken()

Your session is now authorized to make requests.

## Making Requests

Most API methods take a series of optional parameters. These parameters should be sent in an instance of a
package.Parameters struct (where package is files, fileops, account). If you do not wish to send any optional
parameters, just pass nil for that argument. All required parameters are part of the function signature:

        p := dropbox.Parameters{
            Rev: "01b3f45a",
            FileLimit: "50",
        }

        u := dropbox.Uri{
            Root: "sandbox",
            Path: "path/to/my/file.pdf",
        }

        dropbox.GetMetadata(s, u, p) // sup?

Or:

        dropbox.GetMetadata(s, u, nil)

Any irrelevant parameters will be ignored in the request.

## API Documentation:

http://go.pkgdoc.org/github.com/scottferg/Dropbox-Go/dropbox

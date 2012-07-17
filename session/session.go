package session

import (
    "net/http"
    "net/url"
    "fmt"
    "time"
    "bytes"
    "strings"
    "io/ioutil"
)

var (
    BaseApiUrl = "api.dropbox.com"
    BaseWebUrl = "www.dropbox.com"

    ApiVersion = 1
)

const (
    GET = iota
    POST
    PUT
    DELETE
)

type AccessToken struct {
    Key string
    Secret string
}

type Session struct {
    AppKey string
    AppSecret string
    AccessType string
    Token AccessToken
}

type RequestToken struct {
    Key string
    Secret string
}

func buildApiUrl(path string) string {
    return fmt.Sprintf("https://%s/%d/%s", BaseApiUrl, ApiVersion, path)
}

func buildWebUrl(path string) string {
    return fmt.Sprintf("https://%s/%d/%s", BaseWebUrl, ApiVersion, path)
}

func (s *Session) MakeApiRequest(path string, method int) ([]byte, error) {
    var client http.Client

    var m string
    switch method {
    case GET:
        m = "GET"
    case POST:
        m = "POST"
    case PUT:
        m = "PUT"
    case DELETE:
        m = "DELETE"
    }

    req, err := http.NewRequest(m, buildApiUrl(path), nil)

	if err != nil {
		return nil, err
	}

    auth := s.buildAuthHeader()

    req.Header.Set("Authorization", auth)

    resp, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    return body, err
}

func (s *Session) buildAuthHeader() string {
    // https://gist.github.com/1671416
    var buf bytes.Buffer
    buf.WriteString(`OAuth auth_version="1.0", oauth_signature_method="PLAINTEXT"`)
	fmt.Fprintf(&buf, `, oauth_consumer_key="%s"`, url.QueryEscape(s.AppKey))
	fmt.Fprintf(&buf, `, oauth_timestamp="%v"`, time.Now().Unix())

	signed := ""

    if s.Token.Secret != "" {
		signed = url.QueryEscape(s.Token.Secret)
		fmt.Fprintf(&buf, `, oauth_token="%s"`, url.QueryEscape(s.Token.Key))
	}

	fmt.Fprintf(&buf, `, oauth_signature="%s&%s"`, url.QueryEscape(s.AppSecret), signed)
	return buf.String()
}

func (s *Session) ObtainRequestToken(t *AccessToken) (token string, err error) {
    if body, err := s.MakeApiRequest("oauth/request_token", POST); err != nil {
        panic(err.Error())
    } else {
        tokens := strings.Split(string(body), "&")
        t.Secret = strings.Split(tokens[0], "=")[1]
        t.Key = strings.Split(tokens[1], "=")[1]
    }

    return
}

func (s *Session) ObtainAccessToken(t *AccessToken) (token string, err error) {
    if body, err := s.MakeApiRequest("oauth/access_token", POST); err != nil {
        panic(err.Error())
    } else {
        tokens := strings.Split(string(body), "&")
        t.Secret = strings.Split(tokens[0], "=")[1]
        t.Key = strings.Split(tokens[1], "=")[1]
    }

    return
}

func (s *Session) GenerateAuthorizeUrl(requestToken string) string {
    return fmt.Sprintf("%s?oauth_token=%s", buildWebUrl("oauth/authorize"), requestToken)
}

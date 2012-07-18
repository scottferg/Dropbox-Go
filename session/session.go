package session

import (
    "net/http"
    "net/url"
    "fmt"
    "time"
    "bytes"
    "strings"
    "io/ioutil"
    "encoding/json"
)

var (
    BaseApiUrl = "api.dropbox.com"
    BaseContentUrl = "api-content.dropbox.com"
    BaseWebUrl = "www.dropbox.com"

    ApiVersion = 1
)

const (
    GET = "GET"
    POST = "POST"
    PUT = "PUT"
    DELETE = "DELETE"
)

type AuthError struct {
    ErrorText string `json:"error"`
}

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

func buildContentApiUrl(path string) string {
    return fmt.Sprintf("https://%s/%d/%s", BaseContentUrl, ApiVersion, path)
}

func buildWebUrl(path string) string {
    return fmt.Sprintf("https://%s/%d/%s", BaseWebUrl, ApiVersion, path)
}

func (e AuthError) Error() string {
    return e.ErrorText
}

func (s *Session) DoRequest(url string, method string, file []byte) ([]byte, http.Header, error) {
    req, err := http.NewRequest(method, url, nil)

    var client http.Client

	if err != nil {
        fmt.Println(err.Error())
		return nil, nil, err
	}

    auth := s.buildAuthHeader()

    req.Header.Set("Authorization", auth)

    if file != nil {
        closer := ioutil.NopCloser(bytes.NewReader(file))

        req.Body = closer
        req.ContentLength = int64(len(file))
    }

    resp, err := client.Do(req)

    if err != nil {
        fmt.Println(err.Error())
        return nil, nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    return body, resp.Header, err
}

func (s *Session) MakeContentApiRequest(path string, method string) (b []byte, h http.Header, e error) {
    b, h, e = s.DoRequest(buildContentApiUrl(path), method, nil)
    return
}

func (s *Session) MakeApiRequest(path string, method string) (b []byte, h http.Header, e error) {
    b, h, e = s.DoRequest(buildApiUrl(path), method, nil)
    return
}

func (s *Session) MakeUploadRequest(path string, method string, file []byte) (b []byte, h http.Header, e error) {
    b, h, e = s.DoRequest(buildContentApiUrl(path), method, file)
    return
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

func (s *Session) ObtainRequestToken() (token string, err error) {
    if body, _, err := s.MakeApiRequest("oauth/request_token", POST); err != nil {
        panic(err.Error())
    } else {
        tokens := strings.Split(string(body), "&")
        s.Token.Secret = strings.Split(tokens[0], "=")[1]
        s.Token.Key = strings.Split(tokens[1], "=")[1]
    }

    return
}

func (s *Session) ObtainAccessToken() (token string, err error) {
    body, _, err := s.MakeApiRequest("oauth/access_token", POST)
    
    if err != nil {
        return
    }

    var autherror AuthError
    err = json.Unmarshal(body, &autherror)

    if autherror.ErrorText != "" {
        return token, autherror
    }
        
    tokens := strings.Split(string(body), "&")
    s.Token.Secret = strings.Split(tokens[0], "=")[1]
    s.Token.Key = strings.Split(tokens[1], "=")[1]

    return
}

func GenerateAuthorizeUrl(requestToken string) string {
    return fmt.Sprintf("%s?oauth_token=%s", buildWebUrl("oauth/authorize"), requestToken)
}

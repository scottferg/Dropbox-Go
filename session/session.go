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

func (s *Session) DoRequest(url string, method string, file []byte) ([]byte, error) {
    fmt.Println(url)
    req, err := http.NewRequest(method, url, nil)

    var client http.Client

	if err != nil {
		return nil, err
	}

    auth := s.buildAuthHeader()

    req.Header.Set("Authorization", auth)

    if file != nil {
        closer := ioutil.NopCloser(bytes.NewReader(file))

        req.Body = closer
        req.ContentLength = int64(len(file))
    }

    fmt.Println("Blocking")
    resp, err := client.Do(req)
    fmt.Println("Unblocked!")

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    return body, err
}

func (s *Session) MakeContentApiRequest(path string, method string) ([]byte, error) {
    b, e := s.DoRequest(buildContentApiUrl(path), method, nil)
    return b, e
}

func (s *Session) MakeApiRequest(path string, method string) ([]byte, error) {
    b, e := s.DoRequest(buildApiUrl(path), method, nil)
    return b, e
}

func (s *Session) MakeUploadRequest(path string, method string, file []byte) ([]byte, error) {
    b, e := s.DoRequest(buildContentApiUrl(path), method, file)
    return b, e
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
    body, err := s.MakeApiRequest("oauth/access_token", POST)
    
    if err != nil {
        return
    }

    var autherror AuthError
    err = json.Unmarshal(body, &autherror)

    if autherror.ErrorText != "" {
        return token, autherror
    }
        
    tokens := strings.Split(string(body), "&")
    t.Secret = strings.Split(tokens[0], "=")[1]
    t.Key = strings.Split(tokens[1], "=")[1]

    return
}

func (s *Session) GenerateAuthorizeUrl(requestToken string) string {
    return fmt.Sprintf("%s?oauth_token=%s", buildWebUrl("oauth/authorize"), requestToken)
}

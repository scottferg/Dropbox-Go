// Package files provides API methods for managing and uploading
// individual files
package files

import (
    "encoding/json"
    "fmt"
    "github.com/scottferg/Dropbox-Go/session"
    "github.com/scottferg/Dropbox-Go/api"
    "bytes"
)

type FileError struct {
    ErrorText string `json:"error"`
}

type Contents struct {
    Size        string `json:"size"`
    Rev         string `json:"rev"`
    ThumbExists bool   `json:"thumb_exists"`
    Bytes       int    `json:"bytes"`
    Modified    string `json:"modified"`
    ClientMTime string `json:"client_mtime"`
    Path        string `json:"path"`
    IsDir       bool   `json:"is_dir"`
    Icon        string `json:"icon"`
    Root        string `json:"root"`
    MimeType    string `json:"mime_type"`
    Revision    int    `json:"revision"`
}

type Metadata struct {
    Size        string     `json:"size"`
    Hash        string     `json:"hash"`
    Bytes       int        `json:"bytes"`
    ThumbExists bool       `json:"thumb_exists"`
    Rev         string     `json:"rev"`
    Modified    string     `json:"modified"`
    Path        string     `json:"path"`
    IsDir       bool       `json:"is_dir"`
    Icon        string     `json:"icon"`
    Root        string     `json:"root"`
    Contents    []Contents `json:"contents"`
    Revision    int        `json:"revision"`
}

type Revision struct {
    IsDeleted   bool `json:"is_deleted"`
    Revision    int        `json:"revision"`
    Rev         string     `json:"rev"`
    ThumbExists bool       `json:"thumb_exists"`
    Bytes       int        `json:"bytes"`
    Modified    string     `json:"modified"`
    Path        string     `json:"path"`
    IsDir       bool       `json:"is_dir"`
    Icon        string     `json:"icon"`
    Root        string     `json:"root"`
    MimeType    string `json:"mime_type"`
    Size        string     `json:"size"`
}

type ShareUrl struct {
    Url     string `json:"url"`
    Expires string `json:"expires"`
}

type CopyHash struct {
    CopyRef string `json:"copy_ref"`
    Expires string `json:"expires"`
}

func (e FileError) Error() string {
    return e.ErrorText
}

// GetFile retrieves the metadata for the file at the specified path,
// or the metadata for that path.
func GetFile(s session.Session, uri api.Uri) (file []byte, m Metadata, err error) {
    file, header, err := s.MakeContentApiRequest(fmt.Sprintf("files/%s/%s", uri.Root, uri.Path), session.GET)

    if err != nil {
        return
    }

    // File metadata is in header, body is file
    buf := bytes.NewBufferString(header.Get("x-dropbox-metadata"))
    err = json.Unmarshal(buf.Bytes(), &m)

    return
}

// UploadFile uploads the file to the specified path.  The file's metadata is
// returned as a result.
func UploadFile(s session.Session, file []byte, uri api.Uri) (m Metadata, err error) {
    body, _, err := s.MakeUploadRequest(fmt.Sprintf("files_put/%s/%s", uri.Root, uri.Path), session.PUT, file)

    if err != nil {
        return
    }

    var fe FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &m)

    return
}

// GetMetadata returns the metadata for the specified path.
func GetMetadata(s session.Session, uri api.Uri) (m Metadata, err error) {
    body, _, err := s.MakeApiRequest(fmt.Sprintf("metadata/%s/%s", uri.Root, uri.Path), session.GET)

    if err != nil {
        return
    }

    var fe FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &m)

    return
}

func GetRevisions(s session.Session, uri api.Uri) (m []Revision, err error) {
    body, _, err := s.MakeApiRequest(fmt.Sprintf("revisions/%s/%s", uri.Root, uri.Path), session.GET)

    if err != nil {
        return
    }

    var fe FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &m)

    return
}

func RestoreFile(s session.Session, uri api.Uri, rev string) (m Metadata, err error) {
    body, _, err := s.MakeApiRequest(fmt.Sprintf("restore/%s/%s?rev=%s", uri.Root, uri.Path, rev), session.POST)

    if err != nil {
        return
    }

    var fe FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &m)

    return
}

func Search(s session.Session, uri api.Uri, query string) (m []Revision, err error) {
    body, _, err := s.MakeApiRequest(fmt.Sprintf("search/%s/%s?query=", uri.Root, uri.Path, query), session.POST)

    fmt.Println(string(body))
    if err != nil {
        return
    }

    var fe FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &m)

    return
}

func Share(s session.Session, uri api.Uri) (u ShareUrl, err error) {
    body, _, err := s.MakeApiRequest(fmt.Sprintf("shares/%s/%s", uri.Root, uri.Path), session.POST)

    if err != nil {
        return
    }

    var fe FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &u)

    return
}

func Media(s session.Session, uri api.Uri) (u ShareUrl, err error) {
    body, _, err := s.MakeApiRequest(fmt.Sprintf("media/%s/%s", uri.Root, uri.Path), session.POST)

    if err != nil {
        return
    }

    var fe FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &u)

    return
}

func CopyRef(s session.Session, uri api.Uri) (c CopyHash, err error) {
    body, _, err := s.MakeApiRequest(fmt.Sprintf("copy_ref/%s/%s", uri.Root, uri.Path), session.GET)

    if err != nil {
        return
    }

    var fe FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &c)

    return
}

func Thumbnail(s session.Session, uri api.Uri) (file []byte, m Metadata, err error) {
    file, header, err := s.MakeContentApiRequest(fmt.Sprintf("thumbnails/%s/%s", uri.Root, uri.Path), session.GET)

    if err != nil {
        return
    }

    // File metadata is in header, body is file
    buf := bytes.NewBufferString(header.Get("x-dropbox-metadata"))
    err = json.Unmarshal(buf.Bytes(), &m)

    return
}

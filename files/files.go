package files

import (
    "encoding/json"
    "fmt"
    "../session"
)

const (
    RootSandbox = "sandbox"
    RootDropbox = "dropbox"
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

func (e FileError) Error() string {
    return e.ErrorText
}

func GetFile(s session.Session, root string, path string) (m Metadata, err error) {
    body, err := s.MakeContentApiRequest(fmt.Sprintf("files/%s/%s", root, path), session.GET)

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

func UploadFile(s session.Session, file []byte, root string, path string) (m Metadata, err error) {
    body, err := s.MakeUploadRequest(fmt.Sprintf("files_put/%s/%s", root, path), session.PUT, file)

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

func GetMetadata(s session.Session, root string, path string) (m Metadata, err error) {
    body, err := s.MakeApiRequest(fmt.Sprintf("metadata/%s/%s", root, path), session.GET)

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

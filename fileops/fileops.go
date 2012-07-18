package fileops

import (
    "github.com/scottferg/Dropbox-Go/files"
)

func Copy(s session.Session, root string, to_path string, from_path string, locale string, from_copy_ref string) (c files.Contents, err error) {
    body, _, err := s.MakeApiRequest("fileops/copy", session.POST)

    if err != nil {
        return
    }

    var fe files.FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &c)

    return
}

func CreateFolder(s session.Session, uri api.Uri, locale string) (m files.Metadata, err error) {
    body, _, err := s.MakeApiRequest("fileops/create_folder", session.POST)

    if err != nil {
        return
    }

    var fe files.FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &m)

    return
}

func Delete(s session.Session, uri api.Uri, locale string) (m files.Metadata, err error) {
    body, _, err := s.MakeApiRequest("fileops/delete", session.POST)

    if err != nil {
        return
    }

    var fe files.FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &m)

    return
}

func Move(s session.Session, uri api.Uri, root string, from_path string, to_path string, locale string) (m files.Metadata, err error) {
    body, _, err := s.MakeApiRequest("fileops/move", session.POST)

    if err != nil {
        return
    }

    var fe files.FileError
    err = json.Unmarshal(body, &fe)

    if fe.ErrorText != "" {
        err = fe
        return
    }

    err = json.Unmarshal(body, &m)

    return
}

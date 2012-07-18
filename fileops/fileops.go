package fileops

import (
    "github.com/scottferg/Dropbox-Go/files"
)

func Copy(s session.Session, root string, to_path string, from_path string, locale string, from_copy_ref string) (c files.Contents, err error) {
    params := map[string]string {
        "root": root,
        "to_path": to_path,
    }

    if from_path != "" {
        params["from_path"] = from_path
    }

    if locale != "" {
        params["locale"] = locale
    }

    if from_copy_ref != "" {
        params["from_copy_ref"] = from_copy_ref
    }

    body, _, err := s.MakeApiRequest("fileops/copy", params, session.POST)

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
    params := map[string]string {
        "root": uri.Root,
        "path": uri.Path,
    }

    if locale != "" {
        params["locale"] = locale
    }

    body, _, err := s.MakeApiRequest("fileops/create_folder", params, session.POST)

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
    params := map[string]string {
        "root": uri.Root,
        "path": uri.Path,
    }

    if locale != "" {
        params["locale"] = locale
    }

    body, _, err := s.MakeApiRequest("fileops/delete", params, session.POST)

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
    params := map[string]string {
        "root": root,
        "from_path": from_path,
        "to_path": to_path,
    }

    if locale != "" {
        params["locale"] = locale
    }

    body, _, err := s.MakeApiRequest("fileops/move", params, session.POST)

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

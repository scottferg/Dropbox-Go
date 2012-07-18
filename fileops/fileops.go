package fileops

import (
    "github.com/scottferg/Dropbox-Go/files"
)

type Parameters struct {
    Root string
    ToPath string
    FromPath string
    Locale string
    FromCopyRef string
}

func Copy(s session.Session, root string, to_path string, p *Parameters) (c files.Contents, err error) {
    params := map[string]string {
        "root": root,
        "to_path": to_path,
    }

    if p != nil {
        if p.FromPath != "" {
            params["from_path"] = p.FromPath
        }

        if p.Locale != "" {
            params["locale"] = p.Locale
        }

        if p.FromCopyRef != "" {
            params["from_copy_ref"] = p.FromCopyRef
        }
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

func CreateFolder(s session.Session, uri api.Uri, p *Parameters) (m files.Metadata, err error) {
    params := map[string]string {
        "root": uri.Root,
        "path": uri.Path,
    }

    if p != nil {
        if p.Locale != "" {
            params["locale"] = p.Locale
        }
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

func Delete(s session.Session, uri api.Uri, p *Parameters) (m files.Metadata, err error) {
    params := map[string]string {
        "root": uri.Root,
        "path": uri.Path,
    }

    if p != nil {
        if p.Locale != "" {
            params["locale"] = p.Locale
        }
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

func Move(s session.Session, uri api.Uri, root string, from_path string, to_path string, p *Parameters) (m files.Metadata, err error) {
    params := map[string]string {
        "root": root,
        "from_path": from_path,
        "to_path": to_path,
    }

    if p != nil {
        if p.Locale != "" {
            params["locale"] = p.Locale
        }
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

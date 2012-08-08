package dropbox

import (
	"encoding/json"
)

func Copy(s Session, root string, to_path string, p *Parameters) (c Contents, err error) {
	params := map[string]string{
		"root":    root,
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

	body, _, err := s.MakeApiRequest("fileops/copy", params, POST)

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

func CreateFolder(s Session, uri Uri, p *Parameters) (m Metadata, err error) {
	params := map[string]string{
		"root": uri.Root,
		"path": uri.Path,
	}

	if p != nil {
		if p.Locale != "" {
			params["locale"] = p.Locale
		}
	}

	body, _, err := s.MakeApiRequest("fileops/create_folder", params, POST)

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

func Delete(s Session, uri Uri, p *Parameters) (m Metadata, err error) {
	params := map[string]string{
		"root": uri.Root,
		"path": uri.Path,
	}

	if p != nil {
		if p.Locale != "" {
			params["locale"] = p.Locale
		}
	}

	body, _, err := s.MakeApiRequest("fileops/delete", params, POST)

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

func Move(s Session, uri Uri, root string, from_path string, to_path string, p *Parameters) (m Metadata, err error) {
	params := map[string]string{
		"root":      root,
		"from_path": from_path,
		"to_path":   to_path,
	}

	if p != nil {
		if p.Locale != "" {
			params["locale"] = p.Locale
		}
	}

	body, _, err := s.MakeApiRequest("fileops/move", params, POST)

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

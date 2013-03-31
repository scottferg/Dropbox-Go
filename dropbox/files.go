// Package files provides API methods for managing and uploading
// individual files
package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	IsDeleted   bool   `json:"is_deleted"`
	Revision    int    `json:"revision"`
	Rev         string `json:"rev"`
	ThumbExists bool   `json:"thumb_exists"`
	Bytes       int    `json:"bytes"`
	Modified    string `json:"modified"`
	Path        string `json:"path"`
	IsDir       bool   `json:"is_dir"`
	Icon        string `json:"icon"`
	Root        string `json:"root"`
	MimeType    string `json:"mime_type"`
	Size        string `json:"size"`
}

type ShareUrl struct {
	Url     string `json:"url"`
	Expires string `json:"expires"`
}

type CopyHash struct {
	CopyRef string `json:"copy_ref"`
	Expires string `json:"expires"`
}

type DeltaEntry struct {
	Path     string   `json:"0"`
	Metadata Metadata `json:"1"`
}

type Delta struct {
	Reset   bool          `json:"reset"`
	HasMore bool          `json:"has_more"`
	Cursor  string        `json:"cursor"`
	Entries []interface{} `json:"entries"`
}

func (e FileError) Error() string {
	return e.ErrorText
}

func NewMetadata(m map[string]interface{}) Metadata {
	return Metadata{
		Size:        m["size"].(string),
		Bytes:       int(m["bytes"].(float64)),
		ThumbExists: m["thumb_exists"].(bool),
		Rev:         m["rev"].(string),
		Modified:    m["modified"].(string),
		Path:        m["path"].(string),
		IsDir:       m["is_dir"].(bool),
		Icon:        m["icon"].(string),
		Root:        m["root"].(string),
		Revision:    int(m["revision"].(float64)),
	}
}

// GetFile retrieves the metadata for the file at the specified path,
// or the metadata for that path.
func GetFile(s Session, uri Uri, p *Parameters) (file []byte, m Metadata, err error) {
	params := make(map[string]string)

	if p.Rev != "" {
		params["rev"] = p.Rev
	}

	file, header, err := s.MakeContentApiRequest(fmt.Sprintf("files/%s/%s", uri.Root, uri.Path), params, GET)

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
func UploadFile(s Session, file []byte, uri Uri, p *Parameters) (m Metadata, err error) {

	// Upload method requires that all params are sent in the query string, so we'll set them up here rather
	// than letting the set them
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("files_put/%s/%s", uri.Root, uri.Path))

	if p != nil {
		fmt.Fprint(&buf, "?")

		if p.Locale != "" {
			fmt.Fprintf(&buf, "&locale=%s", p.Locale)
		}

		if p.Overwrite != "" {
			fmt.Fprintf(&buf, "&overwrite=%s", p.Overwrite)
		}

		if p.ParentRev != "" {
			fmt.Fprintf(&buf, "&parent_rev=%s", p.ParentRev)
		}
	}

	body, _, err := s.MakeUploadRequest(buf.String(), nil, PUT, file)

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
func GetMetadata(s Session, uri Uri, p *Parameters) (m Metadata, err error) {
	params := make(map[string]string)

	if p != nil {
		if p.FileLimit != "" {
			params["file_limit"] = p.FileLimit
		}

		if p.Hash != "" {
			params["hash"] = p.Hash
		}

		if p.List != "" {
			params["list"] = p.List
		}

		if p.IncludeDeleted != "" {
			params["include_deleted"] = p.IncludeDeleted
		}

		if p.Rev != "" {
			params["rev"] = p.Rev
		}

		if p.Locale != "" {
			params["locale"] = p.Locale
		}
	}

	body, _, err := s.MakeApiRequest(fmt.Sprintf("metadata/%s/%s", uri.Root, uri.Path), params, GET)

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

// Returns file delta information for the application directory (Incomplete)
func GetDelta(s Session, p *Parameters) (d Delta, err error) {
	params := make(map[string]string)

	if p != nil {
		if p.Cursor != "" {
			params["cursor"] = p.Cursor
		}

		if p.Locale != "" {
			params["locale"] = p.Locale
		}
	}

	body, _, err := s.MakeApiRequest("delta", params, POST)

	if err != nil {
		return
	}

	var fe FileError
	err = json.Unmarshal(body, &fe)

	if fe.ErrorText != "" {
		err = fe
		return
	}

	err = json.Unmarshal(body, &d)

	// A bit hacky, but the interface types need to
	// be converted to DeltaEntry types
	for i, v := range d.Entries {
		entry := v.([]interface{})

		md := entry[1].(map[string]interface{})
		d.Entries[i] = DeltaEntry{
			Path:     entry[0].(string),
			Metadata: NewMetadata(md),
		}
	}

	return
}

func GetRevisions(s Session, uri Uri, p *Parameters) (m []Revision, err error) {
	params := make(map[string]string)

	if p != nil {
		if p.RevLimit != "" {
			params["rev_limit"] = p.RevLimit
		}

		if p.Locale != "" {
			params["locale"] = p.Locale
		}
	}

	body, _, err := s.MakeApiRequest(fmt.Sprintf("revisions/%s/%s", uri.Root, uri.Path), params, GET)

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

func RestoreFile(s Session, uri Uri, rev string, p *Parameters) (m Metadata, err error) {
	params := map[string]string{
		"rev": rev,
	}

	if p != nil && p.Locale != "" {
		params["locale"] = p.Locale
	}

	body, _, err := s.MakeApiRequest(fmt.Sprintf("restore/%s/%s", uri.Root, uri.Path), params, POST)

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

func Search(s Session, uri Uri, query string) (m []Revision, err error) {
	body, _, err := s.MakeApiRequest(fmt.Sprintf("search/%s/%s?query=%s", uri.Root, uri.Path, query), nil, POST)

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

func Share(s Session, uri Uri, p *Parameters) (u ShareUrl, err error) {
	params := make(map[string]string)

	if p != nil {
		if p.Locale != "" {
			params["locale"] = p.Locale
		}

		if p.ShortUrl != "" {
			params["short_url"] = p.ShortUrl
		}
	}

	body, _, err := s.MakeApiRequest(fmt.Sprintf("shares/%s/%s", uri.Root, uri.Path), params, POST)

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

func Media(s Session, uri Uri, p *Parameters) (u ShareUrl, err error) {
	params := make(map[string]string)

	if p != nil && p.Locale != "" {
		params["locale"] = p.Locale
	}

	body, _, err := s.MakeApiRequest(fmt.Sprintf("media/%s/%s", uri.Root, uri.Path), params, POST)

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

func CopyRef(s Session, uri Uri) (c CopyHash, err error) {
	body, _, err := s.MakeApiRequest(fmt.Sprintf("copy_ref/%s/%s", uri.Root, uri.Path), nil, GET)

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

func Thumbnail(s Session, uri Uri, p *Parameters) (file []byte, m Metadata, err error) {
	params := make(map[string]string)

	if p != nil {
		if p.Format != "" {
			params["format"] = p.Format
		}

		if p.Size != "" {
			params["size"] = p.Size
		}
	}

	file, header, err := s.MakeContentApiRequest(fmt.Sprintf("thumbnails/%s/%s", uri.Root, uri.Path), params, GET)

	if err != nil {
		return
	}

	// File metadata is in header, body is file
	buf := bytes.NewBufferString(header.Get("x-dropbox-metadata"))
	err = json.Unmarshal(buf.Bytes(), &m)

	return
}

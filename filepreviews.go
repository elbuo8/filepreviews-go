package filepreviews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	FilePreviewsAPI = "https://blimp-previews.herokuapp.com"
)

type FilePreviews struct {
	Client *http.Client
}

type FilePreviewsOptions struct {
	Size     map[string]int
	Metadata []string
}

type FilePreviewsResult struct {
	MetadataURL string `json:"metadata_url"`
	PreviewURL  string `json:"preview_url"`
}

func New() *FilePreviews {
	return &FilePreviews{Client: http.DefaultClient}
}

func (fp *FilePreviews) Generate(urlStr string, opts *FilePreviewsOptions) (*FilePreviewsResult, error) {
	result := &FilePreviewsResult{}
	values := url.Values{}
	values.Set("url", urlStr)
	if opts.Metadata != nil {
		values.Set("metadata", strings.Join(opts.Metadata, ","))
	}
	if opts.Size != nil {
		var geometry string
		if val, ok := opts.Size["width"]; ok {
			geometry += string(val)
		}
		if val, ok := opts.Size["height"]; ok {
			geometry += "x" + string(val)
		}
		values.Set("size", geometry)
	}
	resp, err := http.Get(FilePreviewsAPI + "?" + values.Encode())
	if err != nil {
		return result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("Invalid status code: %v", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil

}

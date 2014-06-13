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
	FilePreviewsAPI = "https://api.filepreviews.io/v1/"
)

type FilePreviews struct {
	Client *http.Client
}

type FilePreviewsOptions struct {
	Size     map[string]int
	Metadata []string
}

type FilePreviewsResult struct {
	Metadata   map[string]interface{} `json:"metadata"`
	PreviewURL string                 `json:"preview_url"`
}

func New() *FilePreviews {
	return &FilePreviews{Client: http.DefaultClient}
}

func (fp *FilePreviews) Generate(urlStr string, opts *FilePreviewsOptions) (*FilePreviewsResult, error) {
	result := &FilePreviewsResult{}
	resp, err := fp.handleRequest(buildFPURL(urlStr, opts))
	var URLs map[string]interface{}
	err = readRequestJSONBody(resp, &URLs)
	if err != nil {
		return result, err
	}
	resp, err = fp.handleRequest(URLs["metadata_url"].(string))
	var metadata map[string]interface{}
	err = readRequestJSONBody(resp, &metadata)
	if err != nil {
		return result, err
	}
	result.Metadata = metadata
	result.PreviewURL = URLs["metadata_url"].(string)
	return result, nil
}

func buildFPURL(urlStr string, opts *FilePreviewsOptions) string {
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
	return FilePreviewsAPI + "?" + values.Encode()
}

func (fp *FilePreviews) handleRequest(urlStr string) (*http.Response, error) {
	resp, err := http.Get(urlStr)
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("Invalid status code: %v", resp.StatusCode)
	}
	return resp, err
}

func readRequestJSONBody(resp *http.Response, result *map[string]interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}

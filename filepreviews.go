package filepreviews

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	FilePreviewsURL = "https://api.filepreviews.io/v2"
)

type FilePreviews struct {
	Client    *http.Client
	APIURL    string
	APIKey    string
	APISecret string
}

type Options struct {
	URL      string                 `json:"url"`
	Sizes    []string               `json:"sizes"`
	Format   string                 `json:"format"`
	Metadata []string               `json:"metadata"`
	Pages    string                 `json:"pages"`
	Data     map[string]interface{} `json:"data"`
	Uploader struct {
		Destination string                 `json:"destination"`
		Headers     map[string]interface{} `json:"headers"`
	} `json:"uploader"`
}

type size struct {
	Height string `json:"height"`
	Width  string `json:"width"`
}

type preview struct {
	Page         int  `json:"page"`
	Size         size `json:"size"`
	Resized      bool `json:"resized"`
	OriginalSize size `json:"original_size"`
}

type FilePreviewsResult struct {
	ID           string    `json:"id"`
	URL          string    `json:"url"`
	Status       string    `json:"status"`
	Preview      preview   `json:"preview"`
	Thumbnails   []preview `json:"thumbnails"`
	OriginalFile struct {
		Name       string                 `json:"name"`
		Size       int                    `json:"size"`
		TotalPages int                    `json:"total_pages"`
		Metadata   map[string]interface{} `json:"metadata"`
		Extension  string                 `json:"extension"`
		Encoding   string                 `json:"encoding"`
		Mimetype   string                 `json:"mimetype"`
		Type       string                 `json:"type"`
	} `json:"original_file"`
	UserData map[string]interface{} `json:"user_data"`
}

func New(apiKey, apiSecret string) *FilePreviews {
	return &FilePreviews{
		Client:    http.DefaultClient,
		APIURL:    FilePreviewsURL,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

func (fp *FilePreviews) Generate(opts *Options) (*FilePreviewsResult, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	return fp.makeRequest("/previews", "POST", bytes.NewReader(body))
}

func (fp *FilePreviews) Retrive(id string) (*FilePreviewsResult, error) {
	return fp.makeRequest("/previews/"+id, "GET", nil)
}

func (fp *FilePreviews) makeRequest(endpoint, method string, body io.Reader) (*FilePreviewsResult, error) {
	result := FilePreviewsResult{}
	r, _ := http.NewRequest(method, fp.APIURL+endpoint, body)
	r.SetBasicAuth(fp.APIKey, fp.APISecret)
	res, err := fp.Client.Do(r)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 201 {
		return nil, errors.New(http.StatusText(res.StatusCode))
	}
	return &result, json.NewDecoder(res.Body).Decode(result)
}

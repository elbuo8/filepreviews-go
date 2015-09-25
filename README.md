# Filepreviews-Go

Go client library for the [FilePreviews.io](http://filepreviews.io/) service. Generate image previews and metadata from almost any kind of file.

## Installation
```bash
$ go get github.com/elbuo8/filepreviews-go
```

### Example code
```go
fp := filepreviews.New()
opts := &filepreviews.FilePreviewsOptions{}
_, err := fp.Generate("http://www.getblimp.com/images/screenshot1.png", opts)
```

#### Options
You can optinally send an options object.
```go
fp := New()
opts := &FilePreviewsOptions{
	Size: map[string]int{
		"width":  50,
		"height": 100,
	},
	Metadata: []string{"all"},
}
_, err := fp.Generate("http://www.getblimp.com/images/screenshot1.png", opts)
```

## MIT

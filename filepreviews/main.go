package main

import (
	"encoding/json"
	"flag"
	"github.com/elbuo8/filepreviews-go"
	"log"
	"strings"
)

func main() {
	var metadata = strings.Split(*flag.String("metadata", "all", "Possible values are 'ocr', 'psd', 'exif', and 'all'."), ",")
	var width = flag.Int("width", 100, "Specifies maximum value of thumbnail width.")
	var height = flag.Int("height", 200, "Specifies maximum value of thumbnail height.")
	var URL = flag.String("url", "", "Specifies the URL location of the file.")

	flag.Parse()
	if *URL == "" {
		flag.PrintDefaults()
		log.Panicln("No file specified")
	}

	options := &filepreviews.FilePreviewsOptions{
		Size: map[string]int{
			"width":  *width,
			"height": *height,
		}, Metadata: metadata,
	}
	fp := filepreviews.New()
	log.Println(*URL)
	data, err := fp.Generate(*URL, options)
	if err != nil {
		log.Panicln(err)
	}
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Panicln(err)
	}
	log.Println(string(result))
}

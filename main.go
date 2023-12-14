package main

import (
	"flag"
	"mime"
	"os"
	"path/filepath"

	"github.com/abiriadev/dataurl/dataurl"
)

func main() {
	flag.Parse()

	file := flag.Arg(0)
	if file == "" {
		panic("provide at least one file")
	}

	h, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	mime := mime.TypeByExtension(filepath.Ext(file))
	if mime == "" {
		panic("unknown file type")
	}

	err = dataurl.ToDataUrl(mime, h, os.Stdout)
	if err != nil {
		panic(err)
	}
}

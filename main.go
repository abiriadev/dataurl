package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"io"
	"mime"
	"os"
	"path/filepath"
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

	err = ToDataUrl(mime, h, os.Stdout)
	if err != nil {
		panic(err)
	}
}

func ToDataUrl(mime string, data io.Reader, out io.Writer) error {
	bout := bufio.NewWriter(out)
	_, err := bout.WriteString("data:")
	if err != nil {
		return err
	}

	_, err = bout.WriteString(mime)
	if err != nil {
		return err
	}

	_, err = bout.WriteString(";base64,")
	if err != nil {
		return err
	}

	w := base64.NewEncoder(base64.StdEncoding, bout)
	_, err = io.Copy(w, data)
	if err != nil {
		return err
	}

	return nil
}

package dataurl

import (
	"bufio"
	"errors"
	"io"
	"mime"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

type Mime string

func MimeCustom(raw string) Mime {
	// TODO: check for validity later?
	return Mime(raw)
}

func MimeFromBuf(buf []byte) Mime {
	return Mime(mimetype.Detect(buf).String())
}

func MimeFromExt(ext string) (Mime, error) {
	mime := mime.TypeByExtension(ext)
	if mime == "" {
		return Mime(""), errors.New("Can't detect MIME type from the file extension")
	}

	return Mime(mime), nil
}

type MediaStream struct {
	// stdin or `os.File` reader
	In *bufio.Reader

	// `In` will be pointing to stdin if `Path` is empty
	Path string
}

func NewMediaStream(in io.Reader, path string) MediaStream {
	return MediaStream{
		In:   bufio.NewReader(in),
		Path: path,
	}
}

func (ms *MediaStream) Mime() (Mime, error) {
	if ms.Path != "" {
		mime, err := MimeFromExt(filepath.Ext(ms.Path))
		if err == nil {
			return mime, nil
		}
	}

	buf, err := ms.In.Peek(256)
	if err != nil {
		return Mime(""), err
	}

	mime := MimeFromBuf(buf)
	return Mime(mime), nil
}

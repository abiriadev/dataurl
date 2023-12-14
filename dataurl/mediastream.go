package dataurl

import (
	"bufio"
	"io"
	"path/filepath"
)

const HEADER_LIM = 256

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

	buf, err := ms.In.Peek(HEADER_LIM)
	if err != nil && err != io.EOF {
		return Mime(""), err
	}

	mime := MimeFromBuf(buf)
	return Mime(mime), nil
}

func (ms *MediaStream) DataUrl(out io.Writer) error {
	mime, err := ms.Mime()
	if err != nil {
		return err
	}

	return ToDataUrl(mime, ms.In, out)
}

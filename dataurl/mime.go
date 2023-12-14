package dataurl

import (
	"errors"
	"mime"

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

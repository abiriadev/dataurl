package dataurl

import "github.com/gabriel-vasile/mimetype"

type Mime string

func MimeCustom(raw string) Mime {
	// TODO: check for validity later?
	return Mime(raw)
}

func MimeFromBuf(buf []byte) Mime {
	return Mime(mimetype.Detect(buf).String())
}

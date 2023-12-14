package dataurl

import (
	"bufio"
	"encoding/base64"
	"io"

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

func ToDataUrl(mime Mime, data io.Reader, out io.Writer) error {
	bout := bufio.NewWriter(out)
	_, err := bout.WriteString("data:")
	if err != nil {
		return err
	}

	_, err = bout.WriteString(string(mime))
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

package dataurl

import (
	"encoding/base64"
	"io"
)

func ToDataUrl(mime Mime, data io.Reader, out io.Writer) error {
	_, err := out.Write([]byte("data:"))
	if err != nil {
		return err
	}

	_, err = out.Write([]byte(mime))
	if err != nil {
		return err
	}

	_, err = out.Write([]byte(";base64,"))
	if err != nil {
		return err
	}

	enc := base64.NewEncoder(base64.StdEncoding, out)
	_, err = io.Copy(enc, data)
	if err != nil {
		return err
	}

	return nil
}

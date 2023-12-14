package dataurl

import (
	"bufio"
	"encoding/base64"
	"io"
)

func ToDataUrl(mime Mime, data io.Reader, out *bufio.Writer) error {
	_, err := out.WriteString("data:")
	if err != nil {
		return err
	}

	_, err = out.WriteString(string(mime))
	if err != nil {
		return err
	}

	_, err = out.WriteString(";base64,")
	if err != nil {
		return err
	}

	w := base64.NewEncoder(base64.StdEncoding, out)
	_, err = io.Copy(w, data)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"os"
	"path/filepath"

	"github.com/abiriadev/dataurl/dataurl"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := (&cli.App{
		Name:  "dataurl",
		Usage: "print dataurl of given data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "mime",
				Usage: "Set MIME type for data",
			},
		},
		Action: func(ctx *cli.Context) error {
			m := dataurl.MimeCustom(ctx.String("mime"))

			file := ctx.Args().Get(0)
			var data *os.File

			if file == "" {
				data = os.Stdin

				var buf [128]byte
				_, err := data.Read(buf[:])
				if err != nil {
					panic(err)
				}

				data.Seek(0, 0)

				m = dataurl.MimeFromBuf(buf[:])
			} else {
				d, err := os.Open(file)
				if err != nil {
					panic(err)
				}
				data = d

				if m == "" {
					// m = mime.TypeByExtension(filepath.Ext(file))
					m, err = dataurl.MimeFromExt(filepath.Ext(file))
					if err != nil {
						panic(err)
					}
				}
			}

			if m == "" {
				panic("unknown file type")
			}

			err := dataurl.ToDataUrl(dataurl.Mime(m), data, os.Stdout)
			if err != nil {
				panic(err)
			}

			return nil
		},
	}).Run(os.Args); err != nil {
		panic(err)
	}
}

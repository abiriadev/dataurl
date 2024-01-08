package main

import (
	"mime"
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
			m := ctx.String("mime")

			file := ctx.Args().Get(0)

			if file == "" {
				panic("provide at least one file")
			}

			h, err := os.Open(file)

			if err != nil {
				panic(err)
			}

			if m == "" {
				m = mime.TypeByExtension(filepath.Ext(file))
			}

			if m == "" {
				panic("unknown file type")
			}

			err = dataurl.ToDataUrl(dataurl.Mime(m), h, os.Stdout)

			if err != nil {
				panic(err)
			}

			return nil
		},
	}).Run(os.Args); err != nil {
		panic(err)
	}
}

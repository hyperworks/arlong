package main

import (
	"bytes"
	"encoding/json"
	"github.com/codegangsta/cli"
	"github.com/plimble/arlong"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "arlong"
	app.Usage = "Genrate Swagger 2.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Value: ".",
			Usage: "Package path to generate",
		},

		cli.StringFlag{
			Name:  "out, o",
			Value: ".",
			Usage: "Output Path",
		},

		cli.StringFlag{
			Name:  "file, f",
			Value: "swagger.json",
			Usage: "Output file name",
		},

		cli.BoolFlag{
			Name:  "pretty-print, P",
			Usage: "Pretty-print (indent) the resulting JSON",
		},
	}

	app.Action = func(c *cli.Context) {
		p := c.String("path")
		parser := arlong.NewParser(p)
		b, err := parser.JSON()
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}

		if c.Bool("pretty-print") {
			buf := &bytes.Buffer{}
			if err = json.Indent(buf, b, "", "  "); err != nil {
				os.Stderr.WriteString(err.Error())
				return
			}

			b = buf.Bytes()
		}

		if err = ioutil.WriteFile(path.Join(c.String("out"), c.String("file")), b, 0644); err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
	}

	app.Run(os.Args)
}

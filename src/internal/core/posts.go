package core

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

type post struct {
	HTML     string // generated HTML
	markdown []byte // original markdown (might have mixed HTML)
	file     string
	os.FileInfo
}

func ParsePost(file string) (*post, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot open post file")
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get post file's stat")
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read post file")
	}

	res := &post{
		file:     file,
		markdown: b,
		HTML:     "",
		FileInfo: stat,
	}
	return res, nil
}

var (
	rgH1 = regexp.MustCompile(`# (.+)`)
)

func (p *post) Convert() error {
	// TODO: implement a proper parser & transpiler here

	buf := string(p.markdown)
	p.HTML = rgH1.ReplaceAllString(buf, "<h1>$1</h1>")
	return nil
}

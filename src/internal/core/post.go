package core

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type post struct {
	file   string
	source []byte
	os.FileInfo
}

type Post struct {
	Content    string
	Title      string
	ModifiedOn string
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
		source:   b,
		FileInfo: stat,
	}
	return res, nil
}

var (
	rgH1 = regexp.MustCompile(`# (.+)`)
)

func (p *post) Convert() (Post, error) {
	// TODO(mark): implement a proper parser & transpiler here

	return Post{
		// TODO(mark): ensure Title is URL friendly and human-readable
		Title:      strings.Replace(p.FileInfo.Name(), ".md", "", -1),
		ModifiedOn: p.FileInfo.ModTime().UTC().Format(time.RFC3339),
		Content:    string(blackfriday.Run(p.source)),
	}, nil
}

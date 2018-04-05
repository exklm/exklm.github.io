package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/exklm/exklm.github.io/src/internal/core"
)

func main() {
	templateFile := flag.String("template", "src/template/index.html", "/path/to/template")
	postsDir := flag.String("posts-dir", "_posts/", "/path/to/posts_dir")
	flag.Parse()

	// parse templates (in template)
	tmpl, err := template.ParseFiles(*templateFile)
	if err != nil {
		log.Fatalf("cannot parse; err= %v", err)
	}

	buf := bytes.NewBuffer(nil)

	filepath.Walk(*postsDir, func(path string, info os.FileInfo, err error) error {
		if path == *postsDir {
			return nil
		}

		name := info.Name()
		baseName := name[:len(name)-len(filepath.Ext(path))]
		modifiedOn := info.ModTime().UTC().Format(time.RFC3339)
		buf.WriteString("\n<a href=\"#\"><h1>")
		buf.WriteString(baseName)
		buf.WriteString("</h1></a>\n")
		buf.WriteString("<time>")
		buf.WriteString(modifiedOn)
		buf.WriteString("</time>\n")

		src, err := core.ParsePost(path)
		if err != nil {
			return err
		}
		post, err := src.Convert(path)
		if err != nil {
			return err
		}

		out, err := os.Create("posts/" + baseName + ".html")
		if err != nil {
			return err
		}
		if err := tmpl.Execute(out, post); err != nil {
			return err
		}
		out.Close()

		return nil
	})

	log.Printf("\n%s\n", buf.String())

	post := core.Post{
		Content:    buf.String(),
		Title:      "bitsgofer",
		ModifiedOn: time.Now().UTC().Format(time.RFC3339),
	}
	out, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer out.Close()
	if err := tmpl.Execute(out, post); err != nil {
		log.Fatalf("%v", err)
	}

	// // read markdown posts
	// // 	-> convert / transpile to HTML
	// // generate other pages + index pages
	// src, err := core.ParsePost(*postFile)
	// if err != nil {
	// 	log.Fatalf("cannot parse src; err= %v", err)
	// }
	// post, err := src.Convert()
	// if err != nil {
	// 	log.Fatalf("cannot convert; err= %v", err)
	// }

	// if err := tmpl.Execute(os.Stdout, post); err != nil {
	// 	log.Fatalf("cannot execute; err= %v", err)
	// }
}

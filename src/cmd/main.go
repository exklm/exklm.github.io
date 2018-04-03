package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/exklm/exklm.github.io/src/internal/core"
)

func main() {
	templateFile := flag.String("template", "src/template/index.html", "/path/to/template")
	postFile := flag.String("post", "_posts/test.md", "/path/to/post")
	flag.Parse()

	// parse templates (in template)
	tmpl, err := template.ParseFiles(*templateFile)
	if err != nil {
		log.Fatalf("cannot parse; err= %v", err)
	}

	// read markdown posts
	// 	-> convert / transpile to HTML
	// generate other pages + index pages
	post, err := core.ParsePost(*postFile)
	if err != nil {
		log.Fatalf("cannot parse post; err= %v", err)
	}
	if err := post.Convert(); err != nil {
		log.Fatalf("cannot convert; err= %v", err)
	}

	if err := tmpl.Execute(os.Stdout, post); err != nil {
		log.Fatalf("cannot execute; err= %v", err)
	}
}

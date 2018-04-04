Static website for blog.bitsgofer.com.

******

# Development

Need to install
- docker: to run nginx locally while doing dev
- automake: Makefile ftw!


******

# Building blocks

- [mini.css](https://minicss.org): Mininum CSS with flexbox
	- Used `mini-default.css` and `mini-dark.css` (from <https://github.com/Chalarangelo/mini.css/tree/master/dist>).
- [prismjs](http://prismjs.com): Syntax highlighting:
	- Used plugins: [Line Numbers](http://prismjs.com/plugins/line-numbers/), [Autolinker](http://prismjs.com/plugins/autolinker/) and [Command Line](http://prismjs.com/plugins/command-line/).
	- Used, with these options 2 themes `tomorrow-night` and `solarized-light`.

******

## Roadmap

- [x] Create an HTML + CSS + basic JS page for writing text + code.
- [x] Create a CLI tool in Go for parsing mardown posts & generate its HTML, based on the template.
- [ ] Modify CLI to generate the `ls -alF` page (list all posts)
- [ ] Modify CLI to generate the category pages + post count
- [ ] Setup build pipeline with Bazel
- [ ] Replace the markdown package with hand-rolled parser

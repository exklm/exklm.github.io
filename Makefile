THEME=dark
_mini_css=mini-dark.css
_prism_css=prism-solarized-light.css
_prism_js=prism.js

ifeq ("${THEME}", "light")
_mini_css=mini-default.css
_prism_css=prism-tomorrow-night.css

endif

all: assets content

clean:
	rm -rf \
		assets/css/* \
		assets/js/* \
		index.html
	touch \
		assets/css/mini.min.css \
		assets/css/prism.min.css \
		assets/js/prism.min.js \
		index.html
.PHONY: clean

assets:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) exklm/postcss \
		-m \
		-u autoprefixer \
		-u cssnano \
		-o assets/css/mini.min.css \
		src/assets/css/${_mini_css}
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) exklm/postcss \
		-m \
		-u autoprefixer \
		-u cssnano \
		-o assets/css/prism.min.css \
		src/assets/css/${_prism_css}
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) exklm/closure-compiler \
		--js_output_file assets/js/site.min.js \
		--js \
		src/assets/js/${_prism_js} > /dev/null 2>&1
.PHONY: assets

content:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) exklm/html-minifier \
		-o index.html \
		--remove-comments \
		--collapse-whitespace \
		src/template/index.html
.PHONY: content

fix-assets-permission:
	sudo chown -R $(shell id -u):$(shell id -g) assets/
.PHONY: fix-assets-permission

run.local:
	docker run --name nginx --rm \
		-v $(shell pwd):/usr/share/nginx/html:ro \
		-p 80:80 \
		nginx:1.13.10
.PHONY: run.local

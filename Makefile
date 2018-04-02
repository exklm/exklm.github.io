all:
all: assets/css/mini.min.css
# all: assets/js/prism.min.js

run.local:
	docker run --name nginx -v $(shell pwd):/usr/share/nginx/html:ro -p 83:80 nginx:1.13.10
.PHONY: run.local

clean:
	rm -rf \
		assets/css/* \
		assets/js/* \
		index.html
.PHONY: clean

THEME=LIGHT
_mini_css=mini-default.css
_prism_css=prism-tomorrow-night.css
_prism_js=prism.js

ifeq ("${THEME}", "DARK")
_mini_css=mini-dark.css
_prism_css=prism-solarized-light.css
endif

assets/css/mini.min.css:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) exklm/postcss \
		-m \
		-u autoprefixer \
		-u cssnano \
		-o assets/css/mini.min.css \
		src/assets/css/${_mini_css}

assets/js/prism.min.js:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) closure-compiler \
		--js \
		src/js/${_prism_js}



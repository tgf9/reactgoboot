js_src = $(shell find frontend/tailwind.config.js frontend/src -name "*.js" -o -name "*.jsx")
go_src = $(shell find backend -name "*.go")

go_ldflags :=
esbuild_flags :=

ifeq ($(strip $(rel)),1)
go_ldflags := -ldflags="-s -w"
esbuild_flags := --minify
endif

.PHONY: all
all: backend/backend tags

frontend/node_modules: frontend/package.json frontend/package-lock.json
	cd frontend && npm install
	touch $@

backend/public/index.js: frontend/node_modules $(js_src)
	cd frontend && ./node_modules/.bin/esbuild src/index.jsx --bundle \
		$(esbuild_flags) --outfile=/tmp/$(notdir $@)
	mv /tmp/$(notdir $@) $@

backend/public/index.html: frontend/src/index.html
	cp $< $@

backend/public/index.css: frontend/src/index.css frontend/src/index.html $(js_src)
	cd frontend && ./node_modules/.bin/tailwindcss \
		--input $(subst frontend/,./,$<) --output /tmp/$(notdir $@)
	mv /tmp/$(notdir $@) $@

backend/backend: backend/go.mod $(go_src)
backend/backend: backend/public/index.html backend/public/index.css backend/public/index.js
	cd backend && go build -o $(notdir $@) $(go_ldflags)

tags: $(js_src) $(go_src)
	echo $^ | tr  " " "\n" | ctags -L -

.PHONY: build-frontend
build-frontend: backend/public/index.html backend/public/index.css backend/public/index.js

.PHONY: run
run: backend/backend
	./$<

.PHONY: clean
clean:
	rm -f backend/backend \
		backend/public/index.html backend/public/index.css backend/public/index.js

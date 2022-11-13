js_src = $(shell find frontend/webpack.config.js frontend/src -name "*.js")
go_src = $(shell find backend -name "*.go")

webpack_mode := development
go_ldflags :=

ifeq ($(strip $(rel)),1)
webpack_mode := production
go_ldflags := -ldflags="-s -w"
endif

frontend/node_modules: frontend/package.json frontend/package-lock.json
	cd frontend && npm install
	touch $@

backend/dist/main.js: frontend/node_modules $(js_src)
	cd frontend && ./node_modules/webpack/bin/webpack.js build --mode $(webpack_mode) --output-path /tmp
	mv /tmp/main.js $@

backend/dist/index.html: frontend/src/index.html
	cp $< $@

backend/backend: backend/go.mod $(go_src) backend/dist/index.html backend/dist/main.js
	cd backend && go build -o $(notdir $@) $(go_ldflags)

.PHONY: build-frontend
build-frontend: backend/dist/main.js backend/dist/index.html

.PHONY: run
run: backend/backend
	./$<

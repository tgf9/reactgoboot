# React + Go

This is an example of how to create a static Go binary with an embedded React
frontend for Linux.

The server has two responsibilities, one to serve the static files and two to
handle API requests. The binary contains all of the needed browser files.  No
server side rendering is done.

The illusion of visiting different URLs is handled by React Router on the
client side. The client also does API requests from the browser.

The client side code is built using esbuild. The output of esbuild is embedded
in the Go binary. Tailwind is used for CSS. That output is also embedded in the
Go binary.

The build instructions are documented in Make.

**Run**

```
$ make backend/backend --jobs=8
./backend/backend
```

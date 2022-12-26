package public

import "embed"

//go:embed index.html index.css index.js
var FS embed.FS

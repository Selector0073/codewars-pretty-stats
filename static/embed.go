// Package static embeds static assets (images, logos, etc.) used at runtime
// by the service handlers, bundled into the binary via go:embed.
package static

import _ "embed"

// CodewarsLogo contains the raw SVG markup of the Codewars logo, embedded
// into the binary so the handler does not need to read from the filesystem.
//
//go:embed codewars.svg
var CodewarsLogo string

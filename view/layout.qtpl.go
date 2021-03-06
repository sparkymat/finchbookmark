// Code generated by qtc from "layout.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/layout.qtpl:1
package view

//line view/layout.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/layout.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/layout.qtpl:1
func StreamLayout(qw422016 *qt422016.Writer, title string, content string) {
//line view/layout.qtpl:1
	qw422016.N().S(`
  <!DOCTYPE html>
  <html>
    <head>
      <title>`)
//line view/layout.qtpl:5
	qw422016.E().S(title)
//line view/layout.qtpl:5
	qw422016.N().S(`</title>
      <meta charset="UTF-8">
      <meta name="description" content="Bookmark management tool">
      <meta name="keywords" content="bookmark,bookmarks">
      <meta name="author" content="sparkymat">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <link rel="stylesheet" type="text/css" href="/css/material-icons.css">
      <link rel="stylesheet" type="text/css" href="/css/materialize.min.css" media="screen,projection">
    </head>
    <body>
      <div class="container-fluid">
        <h3>markit!</h1>
        `)
//line view/layout.qtpl:17
	qw422016.N().S(content)
//line view/layout.qtpl:17
	qw422016.N().S(`
      </div>
      <script type="text/javascript" src="/js/materialize.min.js"></script>
    </body>
  </html>
`)
//line view/layout.qtpl:22
}

//line view/layout.qtpl:22
func WriteLayout(qq422016 qtio422016.Writer, title string, content string) {
//line view/layout.qtpl:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/layout.qtpl:22
	StreamLayout(qw422016, title, content)
//line view/layout.qtpl:22
	qt422016.ReleaseWriter(qw422016)
//line view/layout.qtpl:22
}

//line view/layout.qtpl:22
func Layout(title string, content string) string {
//line view/layout.qtpl:22
	qb422016 := qt422016.AcquireByteBuffer()
//line view/layout.qtpl:22
	WriteLayout(qb422016, title, content)
//line view/layout.qtpl:22
	qs422016 := string(qb422016.B)
//line view/layout.qtpl:22
	qt422016.ReleaseByteBuffer(qb422016)
//line view/layout.qtpl:22
	return qs422016
//line view/layout.qtpl:22
}

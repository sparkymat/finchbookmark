// Code generated by qtc from "text.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/partial/text.qtpl:1
package partial

//line view/partial/text.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/partial/text.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/partial/text.qtpl:1
func StreamSubHeading(qw422016 *qt422016.Writer, title string) {
//line view/partial/text.qtpl:1
	qw422016.N().S(`
  <h3 class="text-3xl text-gray-900 font-light mb-3 mt-3">`)
//line view/partial/text.qtpl:2
	qw422016.E().S(title)
//line view/partial/text.qtpl:2
	qw422016.N().S(`</h3>
`)
//line view/partial/text.qtpl:3
}

//line view/partial/text.qtpl:3
func WriteSubHeading(qq422016 qtio422016.Writer, title string) {
//line view/partial/text.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/partial/text.qtpl:3
	StreamSubHeading(qw422016, title)
//line view/partial/text.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line view/partial/text.qtpl:3
}

//line view/partial/text.qtpl:3
func SubHeading(title string) string {
//line view/partial/text.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line view/partial/text.qtpl:3
	WriteSubHeading(qb422016, title)
//line view/partial/text.qtpl:3
	qs422016 := string(qb422016.B)
//line view/partial/text.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line view/partial/text.qtpl:3
	return qs422016
//line view/partial/text.qtpl:3
}

//line view/partial/text.qtpl:5
func StreamLogo(qw422016 *qt422016.Writer) {
//line view/partial/text.qtpl:5
	qw422016.N().S(`
  <h1 class="text-3xl font-mono mx-4 my-8 text-center">mark_it!</h1>
`)
//line view/partial/text.qtpl:7
}

//line view/partial/text.qtpl:7
func WriteLogo(qq422016 qtio422016.Writer) {
//line view/partial/text.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/partial/text.qtpl:7
	StreamLogo(qw422016)
//line view/partial/text.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line view/partial/text.qtpl:7
}

//line view/partial/text.qtpl:7
func Logo() string {
//line view/partial/text.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line view/partial/text.qtpl:7
	WriteLogo(qb422016)
//line view/partial/text.qtpl:7
	qs422016 := string(qb422016.B)
//line view/partial/text.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line view/partial/text.qtpl:7
	return qs422016
//line view/partial/text.qtpl:7
}

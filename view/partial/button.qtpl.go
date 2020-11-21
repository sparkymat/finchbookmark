// Code generated by qtc from "button.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/partial/button.qtpl:1
package partial

//line view/partial/button.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/partial/button.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/partial/button.qtpl:1
func StreamSubmitButton(qw422016 *qt422016.Writer, label string) {
//line view/partial/button.qtpl:1
	qw422016.N().S(`
  <button type="submit" class="bg-blue-500 hover:bg-blue-600 hover:shadow-inner text-white font-bold font-mono w-full p-4 mt-2 rounded-sm shadow">`)
//line view/partial/button.qtpl:2
	qw422016.E().S(label)
//line view/partial/button.qtpl:2
	qw422016.N().S(`</button>
`)
//line view/partial/button.qtpl:3
}

//line view/partial/button.qtpl:3
func WriteSubmitButton(qq422016 qtio422016.Writer, label string) {
//line view/partial/button.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/partial/button.qtpl:3
	StreamSubmitButton(qw422016, label)
//line view/partial/button.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line view/partial/button.qtpl:3
}

//line view/partial/button.qtpl:3
func SubmitButton(label string) string {
//line view/partial/button.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line view/partial/button.qtpl:3
	WriteSubmitButton(qb422016, label)
//line view/partial/button.qtpl:3
	qs422016 := string(qb422016.B)
//line view/partial/button.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line view/partial/button.qtpl:3
	return qs422016
//line view/partial/button.qtpl:3
}
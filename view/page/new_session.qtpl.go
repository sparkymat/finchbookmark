// Code generated by qtc from "new_session.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/page/new_session.qtpl:1
package page

//line view/page/new_session.qtpl:1
import "github.com/sparkymat/markit/view/partial"

//line view/page/new_session.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/page/new_session.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/page/new_session.qtpl:3
func StreamNewSession(qw422016 *qt422016.Writer) {
//line view/page/new_session.qtpl:3
	qw422016.N().S(`
  <div class="row">
    <div class="col s12 m6 offset-m3">
      `)
//line view/page/new_session.qtpl:6
	partial.StreamCard(qw422016, sessionForm())
//line view/page/new_session.qtpl:6
	qw422016.N().S(`
    </div>
  </div>
`)
//line view/page/new_session.qtpl:9
}

//line view/page/new_session.qtpl:9
func WriteNewSession(qq422016 qtio422016.Writer) {
//line view/page/new_session.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/page/new_session.qtpl:9
	StreamNewSession(qw422016)
//line view/page/new_session.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line view/page/new_session.qtpl:9
}

//line view/page/new_session.qtpl:9
func NewSession() string {
//line view/page/new_session.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
//line view/page/new_session.qtpl:9
	WriteNewSession(qb422016)
//line view/page/new_session.qtpl:9
	qs422016 := string(qb422016.B)
//line view/page/new_session.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
//line view/page/new_session.qtpl:9
	return qs422016
//line view/page/new_session.qtpl:9
}

//line view/page/new_session.qtpl:11
func streamsessionForm(qw422016 *qt422016.Writer) {
//line view/page/new_session.qtpl:11
	qw422016.N().S(`
  <form action="/login" method="POST">
    <h3>Login</h3>
    `)
//line view/page/new_session.qtpl:14
	partial.StreamTextInput(qw422016, "username", "username", "Username")
//line view/page/new_session.qtpl:14
	qw422016.N().S(`
    `)
//line view/page/new_session.qtpl:15
	partial.StreamPasswordInput(qw422016, "password", "password", "Password")
//line view/page/new_session.qtpl:15
	qw422016.N().S(`
    <div class="row">
      `)
//line view/page/new_session.qtpl:17
	partial.StreamSubmitButton(qw422016, "Login")
//line view/page/new_session.qtpl:17
	qw422016.N().S(`
    </div>
  </form>
`)
//line view/page/new_session.qtpl:20
}

//line view/page/new_session.qtpl:20
func writesessionForm(qq422016 qtio422016.Writer) {
//line view/page/new_session.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/page/new_session.qtpl:20
	streamsessionForm(qw422016)
//line view/page/new_session.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line view/page/new_session.qtpl:20
}

//line view/page/new_session.qtpl:20
func sessionForm() string {
//line view/page/new_session.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line view/page/new_session.qtpl:20
	writesessionForm(qb422016)
//line view/page/new_session.qtpl:20
	qs422016 := string(qb422016.B)
//line view/page/new_session.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line view/page/new_session.qtpl:20
	return qs422016
//line view/page/new_session.qtpl:20
}

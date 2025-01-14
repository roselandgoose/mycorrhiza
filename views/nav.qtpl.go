// Code generated by qtc from "nav.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/nav.qtpl:1
package views

//line views/nav.qtpl:1
import "net/http"

//line views/nav.qtpl:2
import "strings"

//line views/nav.qtpl:3
import "github.com/bouncepaw/mycorrhiza/hyphae/backlinks"

//line views/nav.qtpl:4
import "github.com/bouncepaw/mycorrhiza/l18n"

//line views/nav.qtpl:5
import "github.com/bouncepaw/mycorrhiza/user"

//line views/nav.qtpl:6
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/nav.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/nav.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/nav.qtpl:8
func streamhyphaInfoEntry(qw422016 *qt422016.Writer, h *hyphae.Hypha, u *user.User, action, displayText string) {
//line views/nav.qtpl:8
	qw422016.N().S(`
`)
//line views/nav.qtpl:9
	if u.CanProceed(action) {
//line views/nav.qtpl:9
		qw422016.N().S(`
<li class="hypha-info__entry hypha-info__entry_`)
//line views/nav.qtpl:10
		qw422016.E().S(action)
//line views/nav.qtpl:10
		qw422016.N().S(`">
	<a class="hypha-info__link" href="/`)
//line views/nav.qtpl:11
		qw422016.E().S(action)
//line views/nav.qtpl:11
		qw422016.N().S(`/`)
//line views/nav.qtpl:11
		qw422016.E().S(h.Name)
//line views/nav.qtpl:11
		qw422016.N().S(`">`)
//line views/nav.qtpl:11
		qw422016.E().S(displayText)
//line views/nav.qtpl:11
		qw422016.N().S(`</a>
</li>
`)
//line views/nav.qtpl:13
	}
//line views/nav.qtpl:13
	qw422016.N().S(`
`)
//line views/nav.qtpl:14
}

//line views/nav.qtpl:14
func writehyphaInfoEntry(qq422016 qtio422016.Writer, h *hyphae.Hypha, u *user.User, action, displayText string) {
//line views/nav.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:14
	streamhyphaInfoEntry(qw422016, h, u, action, displayText)
//line views/nav.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:14
}

//line views/nav.qtpl:14
func hyphaInfoEntry(h *hyphae.Hypha, u *user.User, action, displayText string) string {
//line views/nav.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:14
	writehyphaInfoEntry(qb422016, h, u, action, displayText)
//line views/nav.qtpl:14
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:14
	return qs422016
//line views/nav.qtpl:14
}

//line views/nav.qtpl:16
func streamhyphaInfo(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha) {
//line views/nav.qtpl:16
	qw422016.N().S(`
`)
//line views/nav.qtpl:18
	u := user.FromRequest(rq)
	lc := l18n.FromRequest(rq)
	backs := backlinks.BacklinksCount(h)

//line views/nav.qtpl:21
	qw422016.N().S(`
<nav class="hypha-info">
	<ul class="hypha-info__list">
		`)
//line views/nav.qtpl:24
	streamhyphaInfoEntry(qw422016, h, u, "history", lc.Get("ui.history_link"))
//line views/nav.qtpl:24
	qw422016.N().S(`
		`)
//line views/nav.qtpl:25
	streamhyphaInfoEntry(qw422016, h, u, "rename-ask", lc.Get("ui.rename_link"))
//line views/nav.qtpl:25
	qw422016.N().S(`
		`)
//line views/nav.qtpl:26
	streamhyphaInfoEntry(qw422016, h, u, "delete-ask", lc.Get("ui.delete_link"))
//line views/nav.qtpl:26
	qw422016.N().S(`
		`)
//line views/nav.qtpl:27
	streamhyphaInfoEntry(qw422016, h, u, "text", lc.Get("ui.text_link"))
//line views/nav.qtpl:27
	qw422016.N().S(`
		`)
//line views/nav.qtpl:28
	streamhyphaInfoEntry(qw422016, h, u, "attachment", lc.Get("ui.attachment_link"))
//line views/nav.qtpl:28
	qw422016.N().S(`
		`)
//line views/nav.qtpl:29
	streamhyphaInfoEntry(qw422016, h, u, "backlinks", lc.GetPlural("ui.backlinks_link", backs))
//line views/nav.qtpl:29
	qw422016.N().S(`
	</ul>
</nav>
`)
//line views/nav.qtpl:32
}

//line views/nav.qtpl:32
func writehyphaInfo(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha) {
//line views/nav.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:32
	streamhyphaInfo(qw422016, rq, h)
//line views/nav.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:32
}

//line views/nav.qtpl:32
func hyphaInfo(rq *http.Request, h *hyphae.Hypha) string {
//line views/nav.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:32
	writehyphaInfo(qb422016, rq, h)
//line views/nav.qtpl:32
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:32
	return qs422016
//line views/nav.qtpl:32
}

//line views/nav.qtpl:34
func streamsiblingHyphaeHTML(qw422016 *qt422016.Writer, siblings string, lc *l18n.Localizer) {
//line views/nav.qtpl:34
	qw422016.N().S(`
<aside class="sibling-hyphae layout-card">
	<h2 class="sibling-hyphae__title layout-card__title">`)
//line views/nav.qtpl:36
	qw422016.E().S(lc.Get("ui.sibling_hyphae"))
//line views/nav.qtpl:36
	qw422016.N().S(`</h2>
	`)
//line views/nav.qtpl:37
	qw422016.N().S(siblings)
//line views/nav.qtpl:37
	qw422016.N().S(`
</aside>
`)
//line views/nav.qtpl:39
}

//line views/nav.qtpl:39
func writesiblingHyphaeHTML(qq422016 qtio422016.Writer, siblings string, lc *l18n.Localizer) {
//line views/nav.qtpl:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:39
	streamsiblingHyphaeHTML(qw422016, siblings, lc)
//line views/nav.qtpl:39
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:39
}

//line views/nav.qtpl:39
func siblingHyphaeHTML(siblings string, lc *l18n.Localizer) string {
//line views/nav.qtpl:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:39
	writesiblingHyphaeHTML(qb422016, siblings, lc)
//line views/nav.qtpl:39
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:39
	return qs422016
//line views/nav.qtpl:39
}

//line views/nav.qtpl:41
func StreamSubhyphaeHTML(qw422016 *qt422016.Writer, subhyphae string, lc *l18n.Localizer) {
//line views/nav.qtpl:41
	qw422016.N().S(`
`)
//line views/nav.qtpl:42
	if strings.TrimSpace(subhyphae) != "" {
//line views/nav.qtpl:42
		qw422016.N().S(`
<section class="subhyphae">
	<h2 class="subhyphae__title">`)
//line views/nav.qtpl:44
		qw422016.E().S(lc.Get("ui.subhyphae"))
//line views/nav.qtpl:44
		qw422016.N().S(`</h2>
	<nav class="subhyphae__nav">
		<ul class="subhyphae__list">
		`)
//line views/nav.qtpl:47
		qw422016.N().S(subhyphae)
//line views/nav.qtpl:47
		qw422016.N().S(`
		</ul>
	</nav>
</section>
`)
//line views/nav.qtpl:51
	}
//line views/nav.qtpl:51
	qw422016.N().S(`
`)
//line views/nav.qtpl:52
}

//line views/nav.qtpl:52
func WriteSubhyphaeHTML(qq422016 qtio422016.Writer, subhyphae string, lc *l18n.Localizer) {
//line views/nav.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:52
	StreamSubhyphaeHTML(qw422016, subhyphae, lc)
//line views/nav.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:52
}

//line views/nav.qtpl:52
func SubhyphaeHTML(subhyphae string, lc *l18n.Localizer) string {
//line views/nav.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:52
	WriteSubhyphaeHTML(qb422016, subhyphae, lc)
//line views/nav.qtpl:52
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:52
	return qs422016
//line views/nav.qtpl:52
}

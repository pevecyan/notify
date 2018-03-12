package notify

import (
	"fmt"
	"io"
	"os"

	"github.com/coredns/coredns/plugin"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

type Notify struct {
	Next plugin.Handler
}

func (e Notify) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	// This function could be simpler. I.e. just fmt.Println("example") here, but we want to show
	// a slightly more complex example as to make this more interesting.
	// Here we wrap the dns.ResponseWriter in a new ResponseWriter and call the next plugin, when the
	// answer comes back, it will print "example".

	// Wrap.
	pw := NewResponsePrinter(w)

	// Call next plugin (if any).
	return plugin.NextOrFailure(e.Name(), e.Next, ctx, pw, r)
}

// Name implements the Handler interface.
func (e Notify) Name() string { return "notify" }

// ResponsePrinter wrap a dns.ResponseWriter and will write example to standard output when
// WriteMsg is called.
type ResponsePrinter struct {
	dns.ResponseWriter
}

// NewResponsePrinter returns ResponseWriter.
func NewResponsePrinter(w dns.ResponseWriter) *ResponsePrinter {
	return &ResponsePrinter{ResponseWriter: w}
}

// WriteMsg calls the underlying ResponseWriter's WriteMsg method and prints "example" to standard
// output.
func (r *ResponsePrinter) WriteMsg(res *dns.Msg) error {
	fmt.Fprintln(out, ex)
	return r.ResponseWriter.WriteMsg(res)
}

// Make out a reference to os.Stdout so we can easily overwrite it for testing.
var out io.Writer = os.Stdout

// What we will print.
const ex = "notify"

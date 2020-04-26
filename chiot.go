package chiot

import (
	"log"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func init() {
	err := caddy.RegisterModule(Chiot{})
	if err != nil {
		log.Fatal(err)
	}
}

// Chiot implements an HTTP handler for responding to WebDAV clients.
type Chiot struct {
	// The root directory out of which to serve files. If
	// not specified, `{http.vars.root}` will be used if
	// set; otherwise, the current directory is assumed.
	// Accepts placeholders.
	Root string `json:"root,omitempty"`

	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (Chiot) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.chiot",
		New: func() caddy.Module { return new(Chiot) },
	}
}

// Provision sets up the module.
func (c *Chiot) Provision(ctx caddy.Context) error {
	c.logger = ctx.Logger(c)

	if c.Root == "" {
		c.Root = "{http.vars.root}"
	}

	return nil
}

func (c Chiot) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	w.Write([]byte(r.RemoteAddr))
	return nil
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*Chiot)(nil)
	_ caddyfile.Unmarshaler       = (*Chiot)(nil)
)

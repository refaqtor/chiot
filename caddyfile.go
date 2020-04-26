package chiot

import (
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("chiot", parseWebdav)
}

// parseWebdav parses the Caddyfile tokens for the webdav directive.
func parseWebdav(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	c := new(Chiot)
	err := c.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// UnmarshalCaddyfile sets up the handler from Caddyfile tokens.
//
//    chiot [<matcher>] {
//        root <path>
//    }
//
func (c *Chiot) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for d.NextBlock(0) {
			switch d.Val() {
			case "root":
				if c.Root != "" {
					return d.Err("root path already specified")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				c.Root = d.Val()
			default:
				return d.Errf("unrecognized subdirective: %s", d.Val())
			}
		}
	}
	return nil
}

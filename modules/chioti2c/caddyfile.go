package chioti2c

import (
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("i2c", parseCaddyfile)
}

// parseCaddyfile parses the Caddyfile tokens for the webdav directive.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	i2c := new(ChiotI2C)
	err := i2c.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}
	return i2c, nil
}

// UnmarshalCaddyfile sets up the handler from Caddyfile tokens.
//
//    chiot [<matcher>] {
//        root <path>
//    }
//
func (i2c *ChiotI2C) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for d.NextBlock(0) {
			switch d.Val() {
			// case "root":
			// 	if i2c.Root != "" {
			// 		return d.Err("root path already specified")
			// 	}
			// 	if !d.NextArg() {
			// 		return d.ArgErr()
			// 	}
			// 	i2c.Root = d.Val()
			// case "i2cslave":
			// 	if i2c.I2CSlave != 0 {
			// 		return d.Err("i2cslave value already specified")
			// 	}
			// 	if !d.NextArg() {
			// 		return d.ArgErr()
			// 	}
			// 	i2c.I2CSlave = d.Val()
			// case "bus":
			// 	if i2c.I2CBus != 0 {
			// 		return d.Err("bus value already specified")
			// 	}
			// 	if !d.NextArg() {
			// 		return d.ArgErr()
			// 	}
			// 	i2c.I2CBus = d.Val()
			default:
				return d.Errf("unrecognized subdirective: %s", d.Val())
			}
		}
	}
	return nil
}

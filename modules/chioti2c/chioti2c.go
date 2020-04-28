package chioti2c

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func init() {
	err := caddy.RegisterModule(ChiotI2C{})
	if err != nil {
		log.Fatal(err)
	}
}

// ChiotI2C implements an HTTP handler
type ChiotI2C struct {
	// The root directory out of which to serve files. If
	// not specified, `{http.vars.root}` will be used if
	// set; otherwise, the current directory is assumed.
	// Accepts placeholders.
	Root string `json:"root,omitempty"`

	//
	I2CSlave int `json:"slave,omitempty"`

	// Depends on Raspberry Pi version (0 or 1)
	I2CBus uint8 `json:"bus,omitempty"`

	//
	I2CAddress uint8 `json:"address,omitempty"`

	//
	I2CRegister uint8 `json:"register,omitempty"`

	//
	logger *zap.Logger

	Rc *os.File
}

// CaddyModule returns the Caddy module information.
func (ChiotI2C) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.i2c",
		New: func() caddy.Module { return new(ChiotI2C) },
	}
}

// Provision sets up the module.
func (i2c *ChiotI2C) Provision(ctx caddy.Context) error {
	i2c.logger = ctx.Logger(i2c)

	if i2c.Root == "" {
		i2c.Root = "{http.vars.root}"
	}
	if i2c.I2CSlave == 0 {
		i2c.I2CSlave = 0x0703
	}
	if i2c.I2CAddress == 0 {
		i2c.I2CAddress = 0x18
	}
	if i2c.I2CRegister == 0 {
		i2c.I2CRegister = 0x05
	}
	i2c.I2CBus = 1
	return nil
}

func (i2c ChiotI2C) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	err := i2c.Connect()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error in Connect: %v\r\n", err)))
		return err
	}
	w.Write([]byte(fmt.Sprint("Connect completed\r\n")))
	//i2c.Rc = rc
	//defer i2c.Close()
	w.Write([]byte(fmt.Sprint("rc and defer close completed\r\n")))
	t, err := i2c.ReadRegU16BE(i2c.I2CRegister)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error in Read: %v\r\n", err)))
		return err
	}
	w.Write([]byte(fmt.Sprint("Read completed\r\n")))
	tc := float32(t&0x0FFF) / float32(16)
	if int(tc)&0x1000 == 1 {
		tc -= 256.0
	}
	w.Write([]byte(fmt.Sprint("tc calc completed\r\n")))
	w.Write([]byte(fmt.Sprintf("Bus: %d C\r\n", i2c.I2CBus)))
	w.Write([]byte(fmt.Sprintf("Register: %d C\r\n", i2c.I2CRegister)))
	w.Write([]byte(fmt.Sprintf("Slave: %d C\r\n", i2c.I2CSlave)))
	w.Write([]byte(fmt.Sprintf("Address: %d C\r\n", i2c.I2CAddress)))

	w.Write([]byte(fmt.Sprintf("Temperature: %f C\r\n", tc)))
	i2c.Close()
	return nil
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*ChiotI2C)(nil)
	_ caddyfile.Unmarshaler       = (*ChiotI2C)(nil)
)

package notify

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"

	"github.com/mholt/caddy"
)

//init registers this plugin 
func init() {
	caddy.RegisterPlugin("notify", caddy.Plugin{
		ServerType: "dns",
		Action: setup
	})
}


func setup(c *caddy.Controller) error {
	c.Next()
	if c.NextArg() {
		//if there is additional arg throw error
		return plugin.Error("notify", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Notify{Next: next}
	})

	return nil
}
package config

import (
	"path/filepath"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
	ginfacades "github.com/goravel/gin/facades"
)

func init() {
	config := facades.Config()
	config.Add("http", map[string]any{
		// HTTP Driver
		"default": "gin",
		// HTTP Drivers
		"drivers": map[string]any{
			"gin": map[string]any{
				// Optional, default is 4096 KB
				"body_limit":   1024 * 1024 * 4,
				"header_limit": 20480,
				"route": func() (route.Route, error) {
					return ginfacades.Route("gin"), nil
				},
			},
		},
		// HTTP URL
		"url": "http://localhost",
		// HTTP Host
		"host": "",
		// HTTP Port
		"port": config.Env("APP_PORT", "8888"),
		// HTTP Entrance
		"entrance": config.Env("APP_ENTRANCE", "/"),
		// HTTPS Configuration
		"tls": map[string]any{
			// HTTPS Host
			"host": "",
			// HTTPS Port
			"port": config.Env("APP_PORT", "8888"),
			// SSL Certificate
			"ssl": map[string]any{
				// ca.pem
				"cert": config.Env("APP_SSL_CERT", filepath.Join(support.RootPath, "storage/ssl.crt")),
				// ca.key
				"key": config.Env("APP_SSL_KEY", filepath.Join(support.RootPath, "storage/ssl.key")),
			},
		},
	})
}

package main

import (
	"GoSvelteBoilerplate/server"
	"GoSvelteBoilerplate/server/config"

	log "github.com/sirupsen/logrus"
)

// TODO: make this a build time variable
var version string = "1.0.0 (dev)"

func main() {
	log.Info("current Version " + version)
	server.Serve(config.HostAndPort())
}

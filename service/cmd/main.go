package main

import (
	log "github.com/jgfranco17/infrasights/core/pkg/logging"
	root "github.com/jgfranco17/infrasights/service/pkg/root"
)

func main() {
	if err := root.Execute(); err != nil {
		log.Error(err.Error())
	}
}

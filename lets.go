package main

import (
	"example/app"

	"github.com/1ets/lets/boot"
)

// Initiate lets engine
func init() {
	app.OnInit()
	boot.OnInit()
}

// Bootstrap applications
func main() {
	boot.OnMain()
}

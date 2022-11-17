package main

import (
	"TProxxy/ui"
)

var (
	TUI ui.TUI = ui.TUI{}
)

func init() {
	TUI.Init()
}

func main() {
	TUI.Start()
}

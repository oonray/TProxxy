package main

import (
	"flag"
	"fmt"
	"github.com/oonray/TProxxy/ui"
)

var (
	TUI ui.TUI = ui.TUI{}
)

func init() {
	TUI.Init()
}

func main() {
	TUI.Sart()
}

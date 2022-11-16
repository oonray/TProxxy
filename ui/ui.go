package ui

import (
	"github.com/rivo/tview"
	log "github.com/sirupsen/logrus"
)

type TUI struct {
	root *tview.Box
}

func (t *TUI) Init() {
	t.root = tview.NewBox().SetBorder(true).SetTitle("Hello World!")
}

func (t *TUI) Start() error {
	return tview.NewApplication().SetRoot(t.root, true).Run()
}

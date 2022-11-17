package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	app  *tview.Application
	root *tview.Grid
}

func (t *TUI) Init() {
	table := tview.NewTable().SetCell(1, 1, tview.NewTableCell("Hello"))
	t.root = tview.NewGrid().
		SetRows(1, 0, 1).
		SetBorders(true).
		AddItem(table, 0, 0, 3, 3, 0, 0, false)

	t.root.Box.SetBackgroundColor(tcell.Color(tcell.ColorBlack)).
		SetBorderColor(tcell.ColorWhite)
	t.app = tview.NewApplication().SetRoot(t.root, true)
}

func (t *TUI) Start() error {
	return t.app.Run()
}

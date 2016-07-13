package main

import (
    "log"

    "github.com/jroimartin/gocui"
)

func main() {
    // Initial
    g := gocui.NewGui()

    if err := g.Init(); err != nil {
        log.Panicln(err)
    }

    defer g.Close()

    // Layout
    g.SetLayout(layout)

    // Keybinding
    if err := keybindings(g); err != nil {
        log.Panicln(err)
    }

    // Base color style
    g.SelBgColor = gocui.ColorGreen
    g.SelFgColor = gocui.ColorBlack
    g.Cursor = true

    // Loop
    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
}

func keybindings(g *gocui.Gui) error {
    return nil
}

func layout(g *gocui.Gui) error {
    return nil
}

package main

import (
    "fmt"
    "log"
    "time"

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
    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quitAction); err != nil {
        return err
    }

    return nil
}

func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()

    if v, err := g.SetView("main", -1, -1, maxX, maxY); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }

        v.Wrap = true

        fmt.Fprintf(v, "Program started at %s\n", time.Now().Format("2006-01-02 15:04:05"))

        if err := g.SetCurrentView("main"); err != nil {
            return err
        }
    }

    return nil
}

func quitAction(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}

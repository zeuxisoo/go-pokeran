package main

import (
    "fmt"
    "log"
    "time"

    "github.com/jroimartin/gocui"
)

var (
    actionMessage = "Action: %s\n"
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
    // Quit: ctrl + c 　
    if err := g.SetKeybinding("main", gocui.KeyCtrlC, gocui.ModNone, quitAction); err != nil {
        return err
    }

    // Up Right Down Left
    if err := g.SetKeybinding("main", gocui.KeyArrowUp, gocui.ModNone, upAction); err != nil {
        return err
    }

    if err := g.SetKeybinding("main", gocui.KeyArrowRight, gocui.ModNone, rightAction); err != nil {
        return err
    }

    if err := g.SetKeybinding("main", gocui.KeyArrowDown, gocui.ModNone, downAction); err != nil {
        return err
    }

    if err := g.SetKeybinding("main", gocui.KeyArrowLeft, gocui.ModNone, leftAction); err != nil {
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

// Actions
func quitAction(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}

func upAction(g *gocui.Gui, v *gocui.View) error {
    fmt.Fprintf(v, actionMessage, "up")

    cursorDown(g, v)

    return nil
}

func rightAction(g *gocui.Gui, v *gocui.View) error {
    fmt.Fprintf(v, actionMessage, "right")

    cursorDown(g, v)

    return nil
}

func downAction(g *gocui.Gui, v *gocui.View) error {
    fmt.Fprintf(v, actionMessage, "down")

    cursorDown(g, v)

    return nil
}

func leftAction(g *gocui.Gui, v *gocui.View) error {
    fmt.Fprintf(v, actionMessage, "left")

    cursorDown(g, v)

    return nil
}

//
func cursorDown(g *gocui.Gui, v *gocui.View) error {
    if v != nil {
        cursorX, cursorY := v.Cursor()

        if err := v.SetCursor(cursorX, cursorY + 1); err != nil {
            originX, originY := v.Origin()

            if err := v.SetOrigin(originX, originY + 1); err != nil {
                return err
            }
        }
    }
    return nil
}

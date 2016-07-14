package main

import (
    "fmt"
    "log"
    "time"
    "io/ioutil"
    "math/rand"
    "strconv"

    "github.com/jroimartin/gocui"
    "github.com/ptrv/go-gpx"
)

var (
    fakeLocationFile = "data/fake-location.gpx"
    fakeLocationLat  = 0.0
    fakeLocationLon  = 0.0

    actionMessage = "Action: %s, Lat: %f, Lon: %f, Move: %f\n"
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

        // current stored gpx
        gp, err := gpx.ParseFile(fakeLocationFile)

        if err != nil {
            return err
        }

        fakeLocationLat = gp.Waypoints[0].Lat
        fakeLocationLon = gp.Waypoints[0].Lon

        fmt.Fprintf(v, "Program started at %s\n", time.Now().Format("2006-01-02 15:04:05"))
        fmt.Fprintf(v, "GPX Lat: %f, Lon: %f\n", fakeLocationLat, fakeLocationLon)

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
    fakeLocationMove := randomMove()
    fakeLocationLat   = fakeLocationLat + fakeLocationMove

    fmt.Fprintf(v, actionMessage, "up", fakeLocationLat, fakeLocationLon, fakeLocationMove)

    updateGPXFile(fakeLocationLat, fakeLocationLon)
    cursorDown(g, v)

    return nil
}

func rightAction(g *gocui.Gui, v *gocui.View) error {
    fakeLocationMove := randomMove()
    fakeLocationLon   = fakeLocationLon + fakeLocationMove

    fmt.Fprintf(v, actionMessage, "right", fakeLocationLat, fakeLocationLon, fakeLocationMove)

    updateGPXFile(fakeLocationLat, fakeLocationLon)
    cursorDown(g, v)

    return nil
}

func downAction(g *gocui.Gui, v *gocui.View) error {
    fakeLocationMove := randomMove()
    fakeLocationLat   = fakeLocationLat - fakeLocationMove

    fmt.Fprintf(v, actionMessage, "down", fakeLocationLat, fakeLocationLon, fakeLocationMove)

    updateGPXFile(fakeLocationLat, fakeLocationLon)
    cursorDown(g, v)

    return nil
}

func leftAction(g *gocui.Gui, v *gocui.View) error {
    fakeLocationMove := randomMove()
    fakeLocationLon   = fakeLocationLon - fakeLocationMove

    fmt.Fprintf(v, actionMessage, "left", fakeLocationLat, fakeLocationLon, fakeLocationMove)

    updateGPXFile(fakeLocationLat, fakeLocationLon)
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

func updateGPXFile(lat float64, lon float64) error {
    g, err := gpx.ParseFile(fakeLocationFile)

    if err != nil {
        return err
    }

    g.Waypoints[0].Lat = lat
    g.Waypoints[0].Lon = lon

    return ioutil.WriteFile(fakeLocationFile, g.ToXML(), 0644)
}

func randomMove() float64 {
    randomFloat   := 20 * rand.Float64()
    randomInteger := int(randomFloat)
    randomString  := strconv.Itoa(250 + randomInteger)

    f, err := strconv.ParseFloat("0.000" + randomString, 64)

    if err != nil {
        log.Panicln(err)
    }

    return f
}

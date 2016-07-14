# Go Pokeran

A controll fake location program for pokeran

#### Usage

Prepare

    Settings > Accessibility > Add iTerm

Install

    export GOPATH=/path/to/go-pokeran

    glide install

Run Xcode and deploy app to device

    open pokern/pokeran.xcodeproj

Run controll program

    go run *.go

#### Problem

Q. osascript is not allowed assistive access

A. Open `Settings` > Enter `Accessibility` > Add `iTerm`

#!/usr/bin/osascript

delay 0.5

activate application "Xcode"

tell application "System Events"
    tell process "Xcode"
        tell menu bar item "Debug" of menu bar 1
            click

            tell menu item "Simulate Location" of menu 1
                click

                click menu item "fake-location" of menu 1
            end tell
        end tell
    end tell
end tell

activate application "iTerm"

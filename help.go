package main

import "fmt"

func showHelp() {
	fmt.Println(helpText)
}

const helpText = `vmux
The terminal multiplexer inspired by i3

USAGE:
    vmux                  Interactive vmux interface
    vmux ls               List session names (has alias 'vmux ps')
    vmux attach <name>    Attach to a session
    vmux detach           Detach from the current session
    vmux new <name>       Create a new session
    vmux kill <name>      Kill a session

SHORTCUTS:
	Alt+N/Alt+Enter   Create new pane
	Alt+Shift+Q       Close pane
	Alt+Shift+F       Make pane fullscreen
	Alt+Shift+Arrow   Move pane
	Alt+Arrow         Move selection
	Alt+/             Toggle search
`

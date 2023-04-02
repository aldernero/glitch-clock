package main

import (
	"flag"
	"github.com/aldernero/glitch-clock/pkg/tui"
)

func main() {
	var localTime, showDate bool
	var dateSeparator string
	flag.BoolVar(&localTime, "local-time", false, "Use local time instead of UTC")
	flag.BoolVar(&showDate, "show-date", true, "Show date")
	flag.StringVar(&dateSeparator, "sep", "-", "Date separator ('-' or '/', '-' by default)")
	flag.Parse()
	if dateSeparator != "-" && dateSeparator != "/" {
		panic("Invalid date separator")
	}
	tui.StartTea(showDate, localTime, dateSeparator)
}

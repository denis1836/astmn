package ui

import (
	color "github.com/fatih/color"
)

// Abbriviations: c - color | t - text 
var (
	cNC = color.New(color.FgWhite).SprintFunc();
	cInfo = color.New(color.FgBlue).SprintFunc();
	cOk = color.New(color.FgGreen).SprintFunc();
	cWarn = color.New(color.FgYellow).SprintFunc();
	cError = color.New(color.FgRed).SprintFunc();
	tBold = color.New(color.Bold).SprintFunc();
)

func PrintInfo(msg string) {
	color.Printf("%s: %s \n", cInfo("info"), msg)
}

func PrintOk(msg string) {
	color.Printf("%s: %s \n", cOk("info"), msg)
}

func PrintWarn(msg string) {
	color.Printf("%s: %s \n", cWarn("info"), msg)
}

func PrintError(msg string) {
	color.Printf("%s: %s \n", cError("info"), msg)
}

func PrintHeader(title string) {
	color.Printf("\n%s\n\n", tBold(title), msg)
}


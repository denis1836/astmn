package ui

import (
	color "github.com/fatih/color"
)

// c - color | t - text 
var (
	cNC = color.New(color.FgWhite).SprintFunc();
	cInfo = color.New(color.FgBlue).SprintFunc();
	cOk = color.New(color.FgGreen).SprintFunc();
	cWarn = color.New(color.FgYellow).SprintFunc();
	cError = color.New(color.FgRed).SprintFunc();
	tBold = color.New(color.Bold).SprintFunc();
)

func pInfo(msg string) {
	color.Printf("%s: %s \n", cInfo("info"), msg)
}

func pOk(msg string) {
	color.Printf("%s: %s \n", cOk("ok"), msg)
}

func pWarn(msg string) {
	color.Printf("%s: %s \n", cWarn("warning"), msg)
}

func pError(msg string) {
	color.Printf("%s: %s \n", cError("error"), msg)
}

func pBold(title string) {
	color.Printf("%s \n", tBold(title), msg)
}


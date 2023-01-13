package utils

import "github.com/fatih/color"

func Success(str string) {
	color.New(color.Bold, color.FgGreen).PrintlnFunc()(str)
}

func Warning(str string) {
	color.New(color.Bold, color.FgYellow).PrintlnFunc()(str)
}

func Error(str string) {
	color.New(color.Bold, color.FgRed).PrintlnFunc()(str)
}

func Info(str string) {
	color.New(color.BlinkSlow, color.FgHiWhite).PrintlnFunc()(str)
}

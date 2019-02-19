package arena

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
)

var NickName = "Unknown"

func Log(message string, a ...interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("[%s] ", NickName))
	buffer.WriteString(message)
	buffer.WriteString("\n")
	fmt.Printf(buffer.String(), a...)
}

func LogNotice(message string, a ...interface{}) {
	color.Set(color.FgWhite)
	defer color.Unset() // Use it in your function
	Log(message, a...)
}

func LogBroadcast(message string, a ...interface{}) {
	color.Set(color.FgHiGreen)
	defer color.Unset() // Use it in your function
	Log(message, a...)
}

func LogOrder(message string, a ...interface{}) {
	color.Set(color.FgHiBlue)
	defer color.Unset() // Use it in your function
	Log(message, a...)
}

func LogDebug(message string, a ...interface{}) {
	color.Set(color.FgCyan)
	defer color.Unset() // Use it in your function
	Log(message, a...)
}

func LogInfo(message string, a ...interface{}) {
	color.Set(color.FgYellow)
	defer color.Unset() // Use it in your function
	Log(message, a...)
}

func LogWarning(message string, a ...interface{}) {
	color.Set(color.FgRed, color.Bold)
	defer color.Unset() // Use it in your function
	Log(message, a...)
}

func LogError(message string, a ...interface{}) {
	color.Set(color.FgHiRed, color.Bold)
	defer color.Unset() // Use it in your function
	Log(message, a...)
}

package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	prefix  = color.New(color.FgHiCyan, color.Bold).SprintFunc()
	timeFmt = color.New(color.FgHiBlack).SprintFunc()
	infoC   = color.New(color.FgHiWhite).SprintfFunc()
	okC     = color.New(color.FgHiGreen).SprintfFunc()
	warnC   = color.New(color.FgHiYellow).SprintfFunc()
	errC    = color.New(color.FgHiRed, color.Bold).SprintfFunc()
)

func stamp() string {
	return timeFmt(time.Now().Format("15:04:05"))
}

func log(emoji, msg string) {
	fmt.Printf("%s %s %s %s\n", stamp(), prefix("ProcPipe"), emoji, msg)
}

func Info(f string, a ...any)    { log("ℹ️ ", infoC(f, a...)) }
func Success(f string, a ...any) { log("✅", okC(f, a...)) }
func Warn(f string, a ...any)    { log("⚠️ ", warnC(f, a...)) }
func Error(f string, a ...any)   { log("❌", errC(f, a...)) }

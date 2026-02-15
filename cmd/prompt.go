package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func prompt(reader *bufio.Reader, label string) string {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	fmt.Printf("%s: ", green(label))
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func promptInt(reader *bufio.Reader, label string) int64 {
	for {
		s := prompt(reader, label)
		n, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return n
		}
		color.Red("  ⚠ Please enter a valid number")
	}
}

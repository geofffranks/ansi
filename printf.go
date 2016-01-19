package ansi

import (
	"fmt"
	"io"
	"regexp"
)

var (
	colors = map[string] string {
		"k": "00;30", // black
		"K": "01;30", // black (BOLD)

		"r": "00;31", // red
		"R": "01;31", // red (BOLD)

		"g": "00;32", // green
		"G": "01;32", // green (BOLD)

		"y": "00;33", // yellow
		"Y": "01;33", // yellow (BOLD)

		"b": "00;34", // blue
		"B": "01;34", // blue (BOLD)

		"m": "00;35", // magenta
		"M": "01;35", // magenta (BOLD)
		"p": "00;35", // magenta
		"P": "01;35", // magenta (BOLD)

		"c": "00;36", // cyan
		"C": "01;36", // cyan (BOLD)

		"w": "00;37", // white
		"W": "01;37", // white (BOLD)
	}

	re = regexp.MustCompile(`%[kKrRgGyYbBmMpPcCwW]{.*?}`)
)

func colorize(s string) string {
	return re.ReplaceAllStringFunc(s, func (m string) string {
		return "\033[" + colors[m[1:2]] + "m" + m[3:len(m)-1] + "\033[00m"
	})
}

func Printf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(colorize(format), a...)
}

func Fprintf(out io.Writer, format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(out, colorize(format), a...)
}

func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(colorize(format), a...)
}

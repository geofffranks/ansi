package ansi

import (
	"testing"
)

func TestColorizer(t *testing.T) {
	var tests = []struct {
		In string
		Out string
	}{
		{"@k{color}", "\033[00;30mcolor\033[00m"},
		{"@K{COLOR}", "\033[01;30mCOLOR\033[00m"},

		{"@r{color}", "\033[00;31mcolor\033[00m"},
		{"@R{COLOR}", "\033[01;31mCOLOR\033[00m"},

		{"@g{color}", "\033[00;32mcolor\033[00m"},
		{"@G{COLOR}", "\033[01;32mCOLOR\033[00m"},

		{"@y{color}", "\033[00;33mcolor\033[00m"},
		{"@Y{COLOR}", "\033[01;33mCOLOR\033[00m"},

		{"@b{color}", "\033[00;34mcolor\033[00m"},
		{"@B{COLOR}", "\033[01;34mCOLOR\033[00m"},

		{"@m{color}", "\033[00;35mcolor\033[00m"},
		{"@p{color}", "\033[00;35mcolor\033[00m"},
		{"@M{COLOR}", "\033[01;35mCOLOR\033[00m"},
		{"@P{COLOR}", "\033[01;35mCOLOR\033[00m"},

		{"@c{color}", "\033[00;36mcolor\033[00m"},
		{"@C{COLOR}", "\033[01;36mCOLOR\033[00m"},

		{"@w{color}", "\033[00;37mcolor\033[00m"},
		{"@W{COLOR}", "\033[01;37mCOLOR\033[00m"},

		{"@k{black} and @r{red}", "\033[00;30mblack\033[00m and \033[00;31mred\033[00m"},
		{"error: @R{%s}", "error: \033[01;31m%s\033[00m"},

		{"@*{RAINBOW}", "\033[01;31mR\033[00m\033[01;33mA\033[00m\033[01;32mI\033[00m\033[01;36mN\033[00m\033[01;34mB\033[0m\033[01;37mO\033[00m\033[01;31mW\033[00m"},

		{"@s@d@l@f", "@s@d@l@f"},
		{"host error: %s", "host error: %s"},
	}

	for _, test := range tests {
		if colorize(test.In) != test.Out {
			t.Errorf("colorize(`%s`) was `%s`, not `%s`", test.In, colorize(test.In), test.Out)
		}
	}
}

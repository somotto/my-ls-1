package options

import (
	"fmt"
	"os"
	"strings"
)

type Options struct {
	LongFormat bool
	Recursive  bool
	ShowHidden bool
	Reverse    bool
	SortByTime bool
	SortBySize bool
	OnePerLine bool
	NoColor    bool
}

func ParseFlags() (Options, []string) {
	var options Options
	args := os.Args[1:]
	var dirs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") && len(arg) > 1 && arg != "--" {
			for _, flag := range arg[1:] {
				switch flag {
				case 'l':
					options.LongFormat = true
				case 'R':
					options.Recursive = true
				case 'a':
					options.ShowHidden = true
				case 'r':
					options.Reverse = true
				case 't':
					options.SortByTime = true
				case 'S':
					options.SortBySize = true
				case '1':
					options.OnePerLine = true
				case 'G':
					options.NoColor = true
				default:
					fmt.Printf("ls: invalid option -- '%c'\n", flag)
					os.Exit(1)
				}
			}
		} else if arg == "--" {
			dirs = append(dirs, args[i+1:]...)
			break
		} else {
			dirs = append(dirs, arg)
		}
	}

	return options, dirs
}
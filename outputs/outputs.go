package outputs

import (
	"fmt"

	"github.com/fatih/color"
)

type ColorOpts struct {
	Color color.Attribute
	Bold  bool
}

func ColorString(opts *ColorOpts, message string, args ...any) string {
	if opts == nil {
		opts = &ColorOpts{
			Color: color.FgWhite,
			Bold:  false,
		}
	}
	baseColorFormat := color.New(opts.Color)
	if opts.Bold {
		baseColorFormat.Add(color.Bold)
	}
	colorFunc := baseColorFormat.SprintFunc()
	return colorFunc(fmt.Sprintf(message, args...))
}

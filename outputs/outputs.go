package outputs

import (
	"fmt"

	"github.com/fatih/color"
)

func ColorString(textColor color.Attribute, bold bool, message string, args ...any) string {
	baseColorFormat := color.New(textColor)
	if bold {
		baseColorFormat.Add(color.Bold)
	}
	colorFunc := baseColorFormat.SprintFunc()
	return colorFunc(fmt.Sprintf(message, args...))
}

package color

import (
	"fmt"
	"strings"
)

var colorsToRotate = []string{
	Red, Green, Yellow, Blue,
}

// Autof automatically colors the arguments given for a specific format.
// Works exactly like fmt.Sprintf.
//
// WARNING: THIS IS AN EXPERIMENTAL FEATURE AND IT MAY BE REMOVED BASED ON USER FEEDBACK.
// IF YOU USE THIS FEATURE, PLEASE PROVIDE FEEDBACK ON https://github.com/TwiN/go-color/discussions/13
//
// Example:
//
//	println(Autof("Hello, %s.", "world"))
func Autof(format string, a ...any) string {
	if len(a) == 0 { // || Reset == ""
		return fmt.Sprintf(format, a...)
	}
	currentColorIndex, maxColorIndex := 0, len(colorsToRotate)-1
	//var coloredFormat string
	var insideSpecifier bool
	formatLength := len(format)
	var lastMergedIndex int
	sb := &strings.Builder{}
	for i, c := range format {
		if format[i] == '%' {
			// If the character has a % before or after it, it's just an escaped %, so ignore it
			if (i+1 < formatLength && format[i+1] == '%') || (i-1 >= 0 && format[i-1] == '%') {
				// %% is an escape sequence for a single %, so we won't color this one
				continue
			}
			// otherwise we're inside a specifier
			insideSpecifier = true
			// and we need to start coloring it
			sb.WriteString(format[lastMergedIndex:i])
			sb.WriteString(colorsToRotate[currentColorIndex])
			sb.WriteRune(c)
			lastMergedIndex = i + 1
			currentColorIndex++
			if currentColorIndex > maxColorIndex {
				currentColorIndex = 0
			}
			// since we know we're inside a specifier, we can skip the next character now
			continue
		}
		if insideSpecifier {
			switch c {
			case 'v', 's', 'd', 'f', 'F', 'g', 'G', 'e', 'E', 'x', 'X', 'p', 'b', 'T', 'o', 'O', 'c', 'q', 'U', 'w':
				insideSpecifier = false
				// this is the end of the specifier, so we'll add the last character of the specifier and stop coloring
				sb.WriteString(format[lastMergedIndex : i+1])
				sb.WriteString(Reset)
				lastMergedIndex = i + 1
			default:
			}
		}
		continue
	}
	if lastMergedIndex < formatLength {
		sb.WriteString(format[lastMergedIndex:])
	}
	return fmt.Sprintf(sb.String(), a...)
}

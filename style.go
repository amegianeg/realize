package core

import (
	"github.com/fatih/color"
	"strings"
)

var (
	//Output writer
	Output = color.Output
	// Red color
	Red = colorBase(color.FgHiRed)
	// Blue color
	Blue = colorBase(color.FgHiBlue)
	// Green color
	Green = colorBase(color.FgHiGreen)
	// Yellow color
	Yellow = colorBase(color.FgHiYellow)
	// Magenta color
	Magenta = colorBase(color.FgHiMagenta)
	// Cyan color
	Cyan = colorBase(color.FgHiCyan)
	// Cyan color
	White = colorBase(color.FgHiWhite)
)

// ColorBase type
type colorBase color.Attribute

// Regular font with a color
func (c colorBase) Regular(a ...interface{}) string {
	return color.New(color.Attribute(c)).Sprint(a...)
}

// Bold font with a color
func (c colorBase) Bold(a ...interface{}) string {
	return color.New(color.Attribute(c), color.Bold).Sprint(a...)
}

// Prefix is a prefix for each message in cli
func Prefix(value string, color colorBase) string {
	return Yellow.Regular("[") + color.Regular(strings.ToUpper(value)) + Yellow.Regular("]")
}

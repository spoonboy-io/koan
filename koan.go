// Package koan provides utility logging functions designed
// to standardise and colorize the logged output
package koan

import (
	"fmt"
	color "github.com/TwiN/go-color"
	"log"
)

// Logger will hold configuration specific to the logger package
type Logger struct {
	DEBUG bool
}

const (
	INFOPREFIX     = "INFO"
	WARNPREFIX     = "WARN"
	ERRPREFIX      = "ERROR"
	FATALERRPREFIX = "FATAL ERROR"
	DEBUGPREFIX    = "DEBUG"
)

var (
	INFOCOLOR  = color.Gray
	WARNCOLOR  = color.Yellow
	ERRCOLOR   = color.Red
	DEBUGCOLOR = color.Green
)

// Info provides logging with an Info prefix
func (*Logger) Info(msg string) {
	log.Println(write(INFOPREFIX, msg, INFOCOLOR))
}

// Warn provides logging with a Warning prefix
func (*Logger) Warn(msg string) {
	log.Println(write(WARNPREFIX, msg, WARNCOLOR))
}

// Debug provides additional logging if DEBUG is set on/true
func (l *Logger) Debug(msg string) {
	if l.DEBUG {
		log.Println(write(DEBUGPREFIX, msg, DEBUGCOLOR))
	}
}

// Error provides logging with an Error prefix
func (*Logger) Error(msg string, err error) {
	errMsg := fmt.Sprintf("%s (%v)", msg, err)
	log.Println(write(ERRPREFIX, errMsg, ERRCOLOR))
}

// FatalError provides logging with a FatalError prefix and terminates execution
func (*Logger) FatalError(msg string, err error) {
	errMsg := fmt.Sprintf("%s (%v)", msg, err)
	log.Fatalln(write(FATALERRPREFIX, errMsg, ERRCOLOR))
}

// utility to construct colorized log output
func write(prefix string, output, outColor string) string {
	out := fmt.Sprintf("%s: %s", prefix, output)
	return color.Colorize(outColor, out)
}

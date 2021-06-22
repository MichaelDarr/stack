package migrate

import (
	"fmt"
	"log"
	"os"
)

// Log is a logger.
type Log struct {
	verbose bool
}

// Printf prints a formatted string into the log.
func (l *Log) Printf(format string, v ...interface{}) {
	if l.verbose {
		log.Printf(format, v...)
	} else {
		fmt.Fprintf(os.Stderr, format, v...)
	}
}

// Verbose returns true if verbosity is enabled.
func (l *Log) Verbose() bool {
	return l.verbose
}

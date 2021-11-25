package inject

import (
	"log"
)

var verbose = false

func logf(format string, v ...interface{}) {
	if !verbose {
		return
	}
	log.Printf(format, v...)
}

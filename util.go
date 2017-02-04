package err

import (
	"log"
)

// PanicErr panic if argument e is not nil
func PanicErr(e error) {
	if e != nil {
		panic(e)
	}
}

// LogPanicErr panic and log if argument e is not nil
func LogPanicErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}

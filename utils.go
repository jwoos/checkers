package checkers

import (
	"log"
)

func handleError(err error) {
	log.Println(err)
}

func min(a int, b int) {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a int, b int) {
	if a > b {
		return a
	} else {
		return b
	}
}

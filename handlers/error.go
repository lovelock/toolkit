package handlers

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func PrintErrorAndGoon(e error) {
	if e != nil {
		fmt.Printf("%+v\n", e)
	}
}

func PrintErrorAndExit(e error) {
	if e != nil {
		fmt.Printf("%+v\n", e)
		os.Exit(1)
	}
}

func LogErrorAndGoon(e error) {
	if e != nil {
		log.Error().Err(e).Send()
	}
}

func LogErrorAndExit(e error) {
	if e != nil {
		log.Error().Err(e).Send()
		os.Exit(1)
	}
}

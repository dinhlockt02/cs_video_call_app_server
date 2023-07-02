package common

import "github.com/rs/zerolog/log"

func Recovery() {
	if err := recover(); err != nil {
		log.Error().Err(err.(error)).Send()
	}
}

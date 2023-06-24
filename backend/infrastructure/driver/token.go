package driver

import (
	"ref/config"
)

type TokenDriver struct {
	*config.Config
}

func NewTokenDriver(conf *config.Config) *TokenDriver {
	return &TokenDriver{
		conf,
	}
}

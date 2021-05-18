package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"

)

type Config interface {
	comfig.Logger
	Serverer
}

type config struct {
	comfig.Logger
	Serverer

	getter kv.Getter
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Serverer:	NewServerer(getter),
		getter:     getter,
	}
}


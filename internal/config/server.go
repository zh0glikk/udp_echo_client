package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Server struct {
	Ip         string   `fig:"ip"`
	Port 	  string	`fig:"port"`
}

type Serverer interface {
	Server() *Server
}

func NewServerer(getter kv.Getter) Serverer {
	return &serverer{
		getter: getter,
	}
}

type serverer struct {
	getter kv.Getter
	once   comfig.Once
}

func (l *serverer) Server() *Server {
	return l.once.Do(func() interface{} {
		var config Server
		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(l.getter, "server")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out server"))
		}


		return &config
	}).(*Server)
}

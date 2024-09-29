package router

import "eCommerce/config"

type Router struct {
	config *config.Config
}

func NewRouter(config *config.Config) (*Router, error) {
	r := &Router{
		config: config,
	}

	return r, nil
}